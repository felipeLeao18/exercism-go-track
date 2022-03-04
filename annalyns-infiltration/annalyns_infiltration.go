package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	var canAttack bool = true
	if knightIsAwake {
		canAttack = false
	}

	return canAttack
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var canSpy bool = false
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		canSpy = true
	}

	return canSpy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var canSignalPrisoner bool = false
	if prisonerIsAwake && !archerIsAwake {
		canSignalPrisoner = true
	}

	return canSignalPrisoner
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var canFreePrisoner bool = false
	if petDogIsPresent && !archerIsAwake {
		canFreePrisoner = true
	}
	if !petDogIsPresent && prisonerIsAwake {
		if !archerIsAwake && !knightIsAwake {
			canFreePrisoner = true
		}
	}
	return canFreePrisoner
}
