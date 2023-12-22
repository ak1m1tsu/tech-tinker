package model

type Component struct {
	Base
	Name          string
	Description   string
	Price         int
	Type          ComponentType
	Configuration *Configuration
}

type Components []Component

type ComponentType uint8

const (
	ComponentType_CPU ComponentType = iota
	ComponentType_Motherboard
	ComponentType_RAM
	ComponentType_GPU
	ComponentType_Storage
	ComponentType_PowerSupply
	ComponentType_Case
	ComponentType_Cooling
	ComponentType_Fan
	ComponentType_CaseFan
	ComponentType_Unknown
)

var componentTypeNames = map[ComponentType]string{
	ComponentType_CPU:         "CPU",
	ComponentType_Motherboard: "Motherboard",
	ComponentType_RAM:         "RAM",
	ComponentType_GPU:         "GPU",
	ComponentType_Storage:     "Storage",
	ComponentType_Case:        "Case",
	ComponentType_PowerSupply: "Power Supply",
	ComponentType_Cooling:     "Cooling",
	ComponentType_Fan:         "Fan",
	ComponentType_CaseFan:     "Case Fan",
	ComponentType_Unknown:     "Unknown",
}

func (c ComponentType) String() string {
	return componentTypeNames[c]
}
