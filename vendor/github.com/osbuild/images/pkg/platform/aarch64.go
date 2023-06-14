package platform

type Aarch64 struct {
	BasePlatform
	UEFIVendor string
}

func (p *Aarch64) GetArch() Arch {
	return ARCH_AARCH64
}

func (p *Aarch64) GetUEFIVendor() string {
	return p.UEFIVendor
}

func (p *Aarch64) GetPackages() []string {
	packages := p.BasePlatform.FirmwarePackages

	if p.UEFIVendor != "" {
		packages = append(packages,
			"dracut-config-generic",
			"efibootmgr",
			"grub2-efi-aa64",
			"grub2-tools",
			"shim-aa64")
	}

	return packages
}

type Aarch64_IoT struct {
	BasePlatform
	UEFIVendor string
	BootFiles  [][2]string
}

func (p *Aarch64_IoT) GetArch() Arch {
	return ARCH_AARCH64
}

func (p *Aarch64_IoT) GetUEFIVendor() string {
	return p.UEFIVendor
}

func (p *Aarch64_IoT) GetPackages() []string {
	packages := p.BasePlatform.FirmwarePackages

	if p.UEFIVendor != "" {
		packages = append(packages,
			"dracut-config-generic",
			"efibootmgr",
			"grub2-efi-aa64",
			"grub2-tools",
			"shim-aa64")
	}

	return packages
}

func (p *Aarch64_IoT) GetBootFiles() [][2]string {
	return p.BootFiles
}
