module github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/channel-sdk

go 1.26

// channel-sdk defines the MarketplaceGateway interface and Registry used by
// all channel-adapter-{channel} services. Add channel-specific dependencies
// only in the respective adapter service, not here.
