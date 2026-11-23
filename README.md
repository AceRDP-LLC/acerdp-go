# AceRDP API client (Go)

Order and manage AceRDP Windows RDP and Linux VPS servers from your own code. Every method is one operation of the [API reference](https://acerdp.io/api-docs); it returns the answer's `data` and throws `AceRdpError` (with `code`, `status` and the request id) when the API refuses.

Create the key in the panel under Settings > API. Pass `idempotency_key` with anything that charges your balance, so a retry never charges twice.

A read, or a change sent with `idempotency_key`, is tried again after `429` or a server error (twice by default, waiting what `Retry-After` says). `WaitUntilReady(ctx, id, interval, timeout)` waits until a new or reinstalled server is ready.

```sh
go get github.com/acerdp-llc/acerdp-go
```

```go
api := acerdp.New(os.Getenv("ACERDP_API_KEY"))
servers, err := api.ListServers(ctx, nil)
_, err = api.SendAPowerAction(ctx, acerdp.Params{"server": id, "action": "restart"})
```

Go 1.21 or newer, the standard library only. A method returns the answer's `data` as `json.RawMessage`; a refusal is an `*acerdp.Error` with `Code`, `Status` and `Message`.

## Methods

- `AddAnSshKey`: Add an SSH key
- `AddExtraIpAddresses`: Add extra IP addresses
- `AddHeavyUse`: Add Heavy use
- `AddRam`: Add RAM
- `AddStorage`: Add storage
- `ConfigureBruteForceProtection`: Configure brute-force protection
- `CreateAServer`: Create a server
- `CreateAWebhook`: Create a webhook
- `DeleteAWebhook`: Delete a webhook
- `DeleteAnExpiredServer`: Delete an expired server
- `DisableIpRestrictions`: Disable IP restrictions
- `ExtendAServer`: Extend a server
- `FixEverything`: Fix everything
- `GetAPlan`: Get a plan
- `GetAServer`: Get a server
- `GetAServersAlerts`: Get a server's alerts
- `GetATicket`: Get a ticket
- `GetATopUp`: Get a top-up
- `GetAnOrder`: Get an order
- `GetBruteForceProtection`: Get brute-force protection
- `GetDiagnosticsState`: Get diagnostics state
- `GetIpRestrictions`: Get IP restrictions
- `GetLocationOptions`: Get location options
- `GetPowerStatus`: Get power status
- `GetProvisioningStatus`: Get provisioning status
- `GetRecoveryStatus`: Get recovery status
- `GetServerCredentials`: Get server credentials
- `GetTransactionDetails`: Get transaction details
- `GetUsageHistory`: Get usage history
- `GetUserBalance`: Get user balance
- `GetUserProfile`: Get user profile
- `ListAddOns`: List add-ons
- `ListAllTransactions`: List all transactions
- `ListDeliveries`: List deliveries
- `ListEvents`: List events
- `ListExtraIpAddresses`: List extra IP addresses
- `ListInvoices`: List invoices
- `ListLocations`: List locations
- `ListNotifications`: List notifications
- `ListOperatingSystems`: List operating systems
- `ListPlans`: List plans
- `ListServers`: List servers
- `ListSshKeys`: List SSH keys
- `ListTickets`: List tickets
- `ListUpgradeOptions`: List upgrade options
- `ListWebhooks`: List webhooks
- `MarkNotificationsAsRead`: Mark notifications as read
- `MoveAServer`: Move a server
- `QuoteANewServer`: Quote a new server
- `QuoteARenewal`: Quote a renewal
- `ReinstallAServer`: Reinstall a server
- `RemoveAnExtraIpAddress`: Remove an extra IP address
- `RemoveAnSshKey`: Remove an SSH key
- `RemoveHeavyUse`: Remove Heavy use
- `RenewServers`: Renew servers
- `ReplaceIpRestrictions`: Replace IP restrictions
- `ResendADelivery`: Resend a delivery
- `ResetThePassword`: Reset the password
- `RunARepair`: Run a repair
- `RunDiagnostics`: Run diagnostics
- `SendAPowerAction`: Send a power action
- `SendATestEvent`: Send a test event
- `ServersDueForRenewal`: Servers due for renewal
- `StartTheRecoveryHour`: Start the recovery hour
- `TopUpYourBalance`: Top up your balance
- `UnblockAnAddress`: Unblock an address
- `UpdateAServer`: Update a server
- `UpgradeAServer`: Upgrade a server
