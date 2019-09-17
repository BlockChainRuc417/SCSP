
package main


import (
	"fmt"
	"errors"
	"sync"
	"time"
)


const (
   twepoch            = int64(1417937700000) // default timestamp 1449473700000
   DistrictIdBits     = uint(5)
   NodeIdBits         = uint(9)
   sequenceBits       = uint(10)

   maxNodeId          = -1 ^ (-1 << NodeIdBits)
   maxDistrictId      = -1 ^ (-1 << DistrictIdBits)

   nodeIdShift        = sequenceBits
   DistrictIdShift    = sequenceBits + NodeIdBits
   timestampLeftShift = sequenceBits + NodeIdBits + DistrictIdBits
   sequenceMask       = -1 ^ (-1 << sequenceBits)
   maxNextIdsNum      = 100
)


type IdWorker struct {
   sequence      int64
   lastTimestamp int64
   nodeId        int64
   twepoch       int64
   districtId    int64
   mutex         sync.Mutex
}


// =================================================================
//        NewIdWorker - new a snowflake id generator object.
// =================================================================
func NewIdWorker(NodeId int64) (*IdWorker, error) {
   var districtId int64
   districtId = 1
   idWorker := &IdWorker{}
   if NodeId > maxNodeId || NodeId < 0 {
      fmt.Sprintf("NodeId Id can't be greater than %d or less than 0", maxNodeId)
      return nil, errors.New(fmt.Sprintf("NodeId Id: %d error", NodeId))
   }
   if districtId > maxDistrictId || districtId < 0 {
      fmt.Sprintf("District Id can't be greater than %d or less than 0", maxDistrictId)
      return nil, errors.New(fmt.Sprintf("District Id: %d error", districtId))
   }
   idWorker.nodeId = NodeId
   idWorker.districtId = districtId
   idWorker.lastTimestamp = -1
   idWorker.sequence = 0
   idWorker.twepoch = twepoch
   idWorker.mutex = sync.Mutex{}
   fmt.Sprintf("worker starting. timestamp left shift %d, District id bits %d, worker id bits %d, sequence bits %d, workerid %d", timestampLeftShift, DistrictIdBits, NodeIdBits, sequenceBits, NodeId)
   return idWorker, nil
}


// =====================================================
//        timeGen - generate a unix millisecond.
// =====================================================
func timeGen() int64 {
   return time.Now().UnixNano() / int64(time.Millisecond)
}


// ==========================================================
//      tilNextMillis - spin wait till next millisecond.
// ==========================================================
func tilNextMillis(lastTimestamp int64) int64 {
   timestamp := timeGen()
   for timestamp <= lastTimestamp {
      timestamp = timeGen()
   }
   return timestamp
}


// ============================================
//      NextId - get a snowflake id.
// ============================================
func (id *IdWorker) NextId() (int64, error) {
   id.mutex.Lock()
   defer id.mutex.Unlock()
   return id.nextid()
}


func (id *IdWorker) nextid() (int64, error) {
   timestamp := timeGen()
   if timestamp < id.lastTimestamp {
      return 0, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", id.lastTimestamp-timestamp))
   }

   if id.lastTimestamp == timestamp {
      id.sequence = (id.sequence + 1) & sequenceMask
      if id.sequence == 0 {
         timestamp = tilNextMillis(id.lastTimestamp)
      }
   } else {
      id.sequence = 0
   }

   id.lastTimestamp = timestamp
   return ((timestamp - id.twepoch) << timestampLeftShift) | (id.districtId << DistrictIdShift) | (id.nodeId << nodeIdShift) | id.sequence, nil
}
