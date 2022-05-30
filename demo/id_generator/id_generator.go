package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/**
 * @Author  jackie.lqj
 * @Date  2022/5/18 19:06
 * @Description 雪花算法生产分布式ID
 */

const (
	// 机器id位数
	workerBits uint8 = 10
	// 序列号位数
	numberBits uint8 = 12
	// 机器id最大值
	workerMax int64 = -1 ^ (-1 << workerBits)
	// 序列号最大值
	numberMax int64 = -1 ^ (-1 << numberBits)
	// 时间戳偏移量
	timeShift uint8 = workerBits + numberBits
	// 机器id偏移量
	workerShift uint8 = numberBits
	// 开始时间，如果在程序跑了一段时间修改了这个值 可能会导致生成相同的ID
	startTime int64 = 1640966400000
)

// Worker 工作节点
type Worker struct {
	// 互斥锁
	mu sync.Mutex
	// 时间戳
	timestamp int64
	// 机器id
	workerId int64
	// 序列号
	number int64
}

// NewWorker 创建工作节点实例
func NewWorker(workerId int64) (*Worker, error) {
	// 校验工作id合法性
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("worker id excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// GetId 生产id
func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	// 同一时间下，序列号增加
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	// 组装ID
	ID := (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}

func main() {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(node.GetId())
	}
}
