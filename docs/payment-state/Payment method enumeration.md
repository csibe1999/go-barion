Payment method enumeration
This enum indicates the method used to attempt the completion of a payment.

API usage
This enum is used by the following API endpoints:
Payment-PaymentState-v4

|  Enum value  | Byte / int value |                                                                     Description                                                                      |
| :----------: | :--------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------: |
|   Unknown    |        0         |                                                       The completion method is not known yet.                                                        |
| BarionWallet |        10        |                       The payment was completed with a funding source (e-money or bank card) from a registered Barion wallet.                        |
|   BankCard   |        20        |                                                     The payment was completed with a bank card.                                                      |
|  GooglePay   |        30        |                                The payment was completed via Google Pay, using a bank card or other means of payment.                                |
|   ApplePay   |        40        |                                The payment was completed via Apple Pay, using a bank card or other means of payment.                                 |
| OpenBanking  |        50        |                                                    The payment was completed via a bank transfer.                                                    |
|    Other     |       200        | The payment was completed by some other method. Note: This is a reserved value for future method implementations, and should not be used explicitly. |
