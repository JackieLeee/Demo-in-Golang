package adapter

import "testing"

/**
 * @Author: Jackie
 * @Date: 2022/8/22
 * @Description: 适配器模式
 */

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
