package device

import (
	"github.com/iniwex5/vohive/internal/config"
)

// DefaultFreeDeviceLimit 表示允许的设备数量上限。0 表示不限量：前端凭
// deviceLimit > 0 判断是否展示配额徽章，为 0 时徽章隐藏，后端各闸门函数
// 也一律放行。
const DefaultFreeDeviceLimit = 0

// FreeDeviceLimitReached 恒为 false：不再限制设备数量。
func FreeDeviceLimitReached(count int) bool {
	return false
}

func FreeDeviceAddLimitMessage() string {
	return ""
}

func FreeDeviceWorkerLimitMessage() string {
	return ""
}

// FreeDeviceLimitAllowsConfiguredDevice 恒为 true：不再限制设备数量。
func FreeDeviceLimitAllowsConfiguredDevice(devices []config.DeviceConfig, deviceID string) bool {
	return true
}
