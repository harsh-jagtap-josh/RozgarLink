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

func MapAddressToEmployer(employer EmployerResponse, address Address) EmployerResponse {
	// Map all address fields to output worker object
	employer.Location = address.ID
	employer.Details = address.Details
	employer.Street = address.Street
	employer.City = address.City
	employer.State = address.State
	employer.Pincode = address.Pincode

	return employer
}

func MatchAddressWorker(address Address, worker WorkerResponse) bool {
	if address.Details == worker.Details && address.Street == worker.Street && address.State == worker.State && address.City == worker.City && address.Pincode == worker.Pincode {
		return true
	}
	return false
}

func MatchAddressEmployer(address Address, employer EmployerResponse) bool {
	if address.Details == employer.Details && address.Street == employer.Street && address.State == employer.State && address.City == employer.City && address.Pincode == employer.Pincode {
		return true
	}
	return false
}
