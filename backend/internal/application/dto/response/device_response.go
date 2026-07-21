package response

type DeviceResponse struct {
	ID   int64
	Name string
}

type CreateDeviceResponse struct {
	Device DeviceResponse
}
