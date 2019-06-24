package message

const (
	// TypeInvalid is type id for invalid type
	TypeInvalid int8 = 0000

	// TypeNone is type id for BaseMessage
	TypeNone int8 = 0001

	// TypeCommand is type id for CommandMessage
	TypeCommand int8 = 0010

	// TypeRoutingConfigure is type id for RoutingConfigureMessage
	TypeRoutingConfigure int8 = 0020

	// TypeClientConfigure is type id for ClientConfigureMessage
	TypeClientConfigure int8 = 0021

	// TypeConsole is type id for Console
	TypeConsole int8 = 0040

	// TypeProxy is type id for Proxy
	TypeProxy int8 = 0050
)

func init() {
	// Temporary handling for backward compatibility

	//msgpack.RegisterExt(TypeNone, &BaseMessage{})
	//msgpack.RegisterExt(TypeCommand, &Command{})
	//msgpack.RegisterExt(TypeRoutingConfigure, &RoutingConfigure{})
	//msgpack.RegisterExt(TypeClientConfigure, &ClientConfigure{})
	//msgpack.RegisterExt(TypeConsole, &Console{})
	//msgpack.RegisterExt(TypeProxy, &Proxy{})
}
