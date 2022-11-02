package nanoarch

// Save writes the current state to the filesystem.
func (na *naEmulator) Save() error {
	na.Lock()
	defer na.Unlock()

	ss, err := getSaveState()
	if err != nil {
		return err
	}
	na.storage.Save(na.GetTimestampedPath(), ss)
	if err := na.storage.Save(na.GetHashPath(), ss); err != nil {
		return err
	}

	if sram := getSaveRAM(); sram != nil {
		if err := na.storage.Save(na.GetSRAMPath(), sram); err != nil {
			return err
		}
	}
	return nil
}

// Load restores the state from the filesystem.
// Deadlock warning: locks the emulator.
func (na *naEmulator) Load() (err error) {
	na.Lock()
	defer na.Unlock()

	if sramState, err := fromFile(na.GetSRAMPath()); err == nil {
		restoreSaveRAM(sramState)
	}
	if saveState, err := fromFile(na.GetHashPath()); err == nil {
		return restoreSaveState(saveState)
	}
	return
}
