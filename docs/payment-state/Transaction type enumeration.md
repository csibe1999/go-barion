Transaction type enumeration
This enum indicates the type of a transaction.

Included in
Transaction types are used in the following structures:
DetailedPaymentTransaction
StatementDownload

|              Enum value               | Payment-related? |                                                                           Description                                                                           |
| :-----------------------------------: | :--------------: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------: |
|              Commission               |        No        |                                                            E-money transfer fee. Always waived (0).                                                             |
|                 Shop                  |       Yes        |                                                      Completed e-money payment between a user and a shop.                                                       |
|        TransferToExistingUser         |        No        |                                                       E-money transfer between two existing Barion users.                                                       |
|      TransferToTechnicalAccount       |        No        |                                         E-money transfer from an existing Barion user to a currently non-existent user.                                         |
|     TransferFromTechnicalAccount      |        No        |                                    E-money credited to a Barion Wallet that wasn't yet created at the time of the transfer.                                     |
|                Storno                 |        No        |                 E-money credited back to the source Barion Wallet after the target, non-existent Barion Wallet hasn't been created to claim it.                 |
|              WithdrawFee              |        No        |                                                 Fee of a bank transfer to a bank account in the same currency.                                                  |
|              ClosureFee               |        No        |                                                                 Fee of closing a Barion Wallet.                                                                 |
|         StornoBankTransferFee         |        No        |        Fee of canceling a Barion Wallet-to-bank account transfer when declined by the bank holding Barion's transaction account in the target currency.         |
|            CashDepositFee             |        No        |                                             Fee charged for topping up a Barion Wallet using cash at a bank branch.                                             |
|           ForeignDepositFee           |        No        |              Fee incurred for topping up a Barion Wallet from a bank account held in a country where Barion doesn't have a domestic bank account.               |
|          ForeignWithdrawFee           |        No        |        Fee incurred when transferring money from a Barion Wallet to a bank account held in a country where Barion doesn't have a domestic bank account.         |
|           ForeignClosureFee           |        No        |     Fee incurred when a Barion Wallet is closed and emptied into a bank account that's held in a country where Barion doesn't have a domestic bank account.     |
|                Reserve                |       Yes        |                                                               Reserved amount in a Barion Wallet.                                                               |
|             StornoReserve             |       Yes        |                               The remainder amount returned when a previously reserved payment was finished with a lower amount.                                |
|             CardTopUpFee              |        No        |                                                      Fee of topping up a Barion Wallet using a bank card.                                                       |
|           CardProcessingFee           |       Yes        |                                                       Bank card processing fee incurred by a Barion shop.                                                       |
|              GatewayFee               |       Yes        |                                                     The Barion Smart Gateway fee incurred by a Barion shop.                                                     |
|        CardProcessingFeeStorno        |       Yes        |                          The previously reserved bank card processing fee was refunded to a Barion shop after a Barion Wallet payment.                          |
|              Unspecified              |      No/Yes      |                                             Placeholder transaction when the transaction type isn't yet determined.                                             |
|                  In                   |        No        |                                             The amount credited to a Barion Wallet in a top-up using bank transfer.                                             |
|               Withdraw                |        No        |                                      The amount debited from a Barion Wallet when money is transferred to a bank account.                                       |
|              CashDeposit              |        No        |                                                    The amount credited to a Barion Wallet in a cash top-up.                                                     |
|            ForeignDeposit             |        No        |             The amount credited to a Barion Wallet from a bank transfer from a bank in a country where Barion doesn't have a domestic bank account.             |
|        ForeignBankTransferFee         |        No        |                Fee charged after a Barion Wallet top-up originating from a bank in a country where Barion doesn't have a domestic bank account.                 |
|          CashBankTransferFee          |        No        |                                  The fee for handling a manual cash top-up to a non-existent (usually mistyped) Barion Wallet.                                  |
|              CustodyFee               |        No        |                                                            Monthly fee for keeping money in custody.                                                            |
|            ForeignWithdraw            |        No        |                   Amount transferred from a Barion Wallet into a bank account in a country where Barion doesn't have a domestic bank account.                   |
|             TransferBack              |        No        |                             The amount transferred back to a bank account when the recipient Barion Wallet couldn't be determined.                              |
|               CardTopUp               |        No        |                                                The amount credited to a Barion Wallet during a bank card top-up.                                                |
|           CardTopUpBankFee            |       Yes        |                                            Fee incurred by Barion to a card acquirer during a bank card transaction.                                            |
|       EmoneySubstractionRestore       |       Yes        |                                         Refund of card acquirer fee when a transaction using a bank card falls through.                                         |
|              CardPayment              |       Yes        |                                                          Bank card payment between a user and a shop.                                                           |
|                Refund                 |       Yes        |                                                             Refund of a payment to a Barion Wallet.                                                             |
|           RefundToBankCard            |       Yes        |                                                               Refund of a payment to a bank card.                                                               |
|  StornoUnSuccessfulRefundToBankCard   |       Yes        |                                             Cancellation of an unsuccessful refund to a bank card to a Barion shop.                                             |
|              UnderReview              |       Yes        |                                                                 Payment is under investigation.                                                                 |
|             ReleaseReview             |       Yes        |                                                             Payment is released from investigation.                                                             |
|          BankTransferPayment          |       Yes        |                                       Bank transfer payment between customer and shop. Used in payment button scenarios.                                        |
|          RefundToBankAccount          |       Yes        |                                            Refund of a payment to a bank account. Used in payment button scenarios.                                             |
| StornoUnSuccessfulRefundToBankAccount |       Yes        |                                   Cancellation of an unsuccessful refund to a bank account. Used in payment button scenarios.                                   |
|        BankTransferPaymentFee         |       Yes        |                                The fee incurred by a Barion shop for a bank transfer payment. Used in payment button scenarios.                                 |
|      BarionBalanceProcessingFee       |       Yes        | The Barion Wallet processing fee incurred by a Barion shop when a payment was made with a Barion Wallet, and the shop was configured to pay the processing fee. |
