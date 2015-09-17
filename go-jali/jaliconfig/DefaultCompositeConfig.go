package jaliconfig

const DefaultCompositeConfig = CompositeConfig{
	Configurators: []Configurator{
		NewJsonConfiguration(&DefaultJsonConfig),
		// TODO: Complete implementation of CompositeConfiguration
		//		NewEnvirionmentConfiguration(DefaultEnvironmentConfig),
		//		NewCommandLineConfiguration(DefaultCommandLineConfig),
	},
}
