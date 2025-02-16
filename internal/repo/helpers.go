package repo

func MapAddressToWorker(workerWithAddress WorkerResponse, address Address) WorkerResponse {
	// Map all address fields to output worker object
	workerWithAddress.Location = address.ID
	workerWithAddress.Details = address.Details
	workerWithAddress.Street = address.Street
	workerWithAddress.City = address.City
	workerWithAddress.State = address.State
	workerWithAddress.Pincode = address.Pincode

	return workerWithAddress
}

func MatchAddress(address Address, worker WorkerResponse) bool {
	if address.Details == worker.Details && address.Street == worker.Street && address.State == worker.State && address.City == worker.City && address.Pincode == worker.Pincode {
		return true
	}
	return false
}
