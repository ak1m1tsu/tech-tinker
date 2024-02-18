package model

type Component struct {
	Base
	Name        string        `db:"name"`
	Description string        `db:"description"`
	Price       int           `db:"price"`
	Type        ComponentType `db:"type"`

	ConfigurationID string `db:"configuration_id"`
	Configuration   *Configuration
}

type Components []Component

type ComponentType uint8

const (
	ComponentTypeCPU ComponentType = iota
	ComponentTypeMotherboard
	ComponentTypeRAM
	ComponentTypeGPU
	ComponentTypeStorage
	ComponentTypePowerSupply
	ComponentTypeCase
	ComponentTypeCooling
	ComponentTypeFan
	ComponentTypeCaseFan
)

var componentTypeNames = map[ComponentType]string{
	ComponentTypeCPU:         "CPU",
	ComponentTypeMotherboard: "Motherboard",
	ComponentTypeRAM:         "RAM",
	ComponentTypeGPU:         "GPU",
	ComponentTypeStorage:     "Storage",
	ComponentTypeCase:        "Case",
	ComponentTypePowerSupply: "Power Supply",
	ComponentTypeCooling:     "Cooling",
	ComponentTypeFan:         "Fan",
	ComponentTypeCaseFan:     "Case Fan",
}

func (c ComponentType) String() string {
	return componentTypeNames[c]
}
