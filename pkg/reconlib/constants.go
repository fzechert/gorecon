/*
This file contains ants used by the reconlib package
*/

package reconlib

const (
	vendorID          = 0x0C45
	productID         = 0x7100
	libusbLogDebug    = 4
	libusbLogNone     = 0
	gousbParallelRead = 2

	controlTransfer byte = 0x09

	txCurrentChannel byte = 0x10

	txSetDisplayChannel  byte = 0x20
	txSetDisplayChannel0 byte = txSetDisplayChannel
	txSetDisplayChannel1 byte = txSetDisplayChannel + 0x01
	txSetDisplayChannel2 byte = txSetDisplayChannel + 0x02
	txSetDisplayChannel3 byte = txSetDisplayChannel + 0x03
	txSetDisplayChannel4 byte = txSetDisplayChannel + 0x04

	txTemperatureAndSpeed  byte = 0x30
	txTemperatureAndSpeed0 byte = txTemperatureAndSpeed
	txTemperatureAndSpeed1 byte = txTemperatureAndSpeed + 0x01
	txTemperatureAndSpeed2 byte = txTemperatureAndSpeed + 0x02
	txTemperatureAndSpeed3 byte = txTemperatureAndSpeed + 0x03
	txTemperatureAndSpeed4 byte = txTemperatureAndSpeed + 0x04

	txGetDeviceSettings byte = 0x50
	txSetDeviceSettings byte = 0x60

	txGetAlarmAndSpeed  byte = 0x70
	txGetAlarmAndSpeed0 byte = txGetAlarmAndSpeed
	txGetAlarmAndSpeed1 byte = txGetAlarmAndSpeed + 0x01
	txGetAlarmAndSpeed2 byte = txGetAlarmAndSpeed + 0x02
	txGetAlarmAndSpeed3 byte = txGetAlarmAndSpeed + 0x03
	txGetAlarmAndSpeed4 byte = txGetAlarmAndSpeed + 0x04

	txSetAlarmAndSpeed  byte = 0x80
	txSetAlarmAndSpeed0 byte = txSetAlarmAndSpeed
	txSetAlarmAndSpeed1 byte = txSetAlarmAndSpeed + 0x01
	txSetAlarmAndSpeed2 byte = txSetAlarmAndSpeed + 0x02
	txSetAlarmAndSpeed3 byte = txSetAlarmAndSpeed + 0x03
	txSetAlarmAndSpeed4 byte = txSetAlarmAndSpeed + 0x04

	rxCurrentChannel  byte = 0x20
	rxCurrentChannel0 byte = rxCurrentChannel
	rxCurrentChannel1 byte = rxCurrentChannel + 0x01
	rxCurrentChannel2 byte = rxCurrentChannel + 0x02
	rxCurrentChannel3 byte = rxCurrentChannel + 0x03
	rxCurrentChannel4 byte = rxCurrentChannel + 0x04

	rxTemperatureAndSpeed  byte = 0x40
	rxTemperatureAndSpeed0 byte = rxTemperatureAndSpeed
	rxTemperatureAndSpeed1 byte = rxTemperatureAndSpeed + 0x01
	rxTemperatureAndSpeed2 byte = rxTemperatureAndSpeed + 0x02
	rxTemperatureAndSpeed3 byte = rxTemperatureAndSpeed + 0x03
	rxTemperatureAndSpeed4 byte = rxTemperatureAndSpeed + 0x04

	rxGetDeviceSettings byte = 0x60

	rxAlarmAndSpeed  byte = 0x80
	rxAlarmAndSpeed0 byte = rxAlarmAndSpeed
	rxAlarmAndSpeed1 byte = rxAlarmAndSpeed + 0x01
	rxAlarmAndSpeed2 byte = rxAlarmAndSpeed + 0x02
	rxAlarmAndSpeed3 byte = rxAlarmAndSpeed + 0x03
	rxAlarmAndSpeed4 byte = rxAlarmAndSpeed + 0x04

	rxAck  byte = 0xF0
	rxNack byte = 0xFA
)
