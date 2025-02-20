package repo

func MapAddressToWorker(workerWithAddress Worker, address Address) Worker {
	// Map all address fields to output worker object
	workerWithAddress.Location = address.ID
	workerWithAddress.Details = address.Details
	workerWithAddress.Street = address.Street
	workerWithAddress.City = address.City
	workerWithAddress.State = address.State
	workerWithAddress.Pincode = address.Pincode

	return workerWithAddress
}

func MapAddressToEmployer(employer Employer, address Address) Employer {
	// Map all address fields to output employer object
	employer.Location = address.ID
	employer.Details = address.Details
	employer.Street = address.Street
	employer.City = address.City
	employer.State = address.State
	employer.Pincode = address.Pincode

	return employer
}

func MapAddressToJob(job Job, address Address) Job {
	// Map all address fields to output job object
	job.Location = address.ID
	job.Details = address.Details
	job.Street = address.Street
	job.City = address.City
	job.State = address.State
	job.Pincode = address.Pincode

	return job
}

func MapAddressToApplication(application Application, address Address) Application {
	// Map all address fields to output application object
	application.PickUpLocation = address.ID
	application.Details = address.Details
	application.Street = address.Street
	application.City = address.City
	application.State = address.State
	application.Pincode = address.Pincode

	return application
}

func MatchAddressWorker(address Address, worker Worker) bool {
	if address.Details == worker.Details && address.Street == worker.Street && address.State == worker.State && address.City == worker.City && address.Pincode == worker.Pincode {
		return true
	}
	return false
}

func MatchAddressEmployer(address Address, employer Employer) bool {
	if address.Details == employer.Details && address.Street == employer.Street && address.State == employer.State && address.City == employer.City && address.Pincode == employer.Pincode {
		return true
	}
	return false
}

func MatchAddressJob(address Address, job Job) bool {
	if address.Details == job.Details && address.Street == job.Street && address.State == job.State && address.City == job.City && address.Pincode == job.Pincode {
		return true
	}
	return false
}

func MatchAddressApplication(address Address, application Application) bool {
	if address.Details == application.Details && address.Street == application.Street && address.State == application.State && address.City == application.City && address.Pincode == application.Pincode {
		return true
	}
	return false
}
