Transaction status enumeration
This enum indicates the current status of a transaction.

Included in
Transaction status is used in the following structures:
DetailedPaymentTransaction
RefundedTransaction
ProcessedTransaction

|            Enum value             | Byte / int value |                                       Description                                       |
| :-------------------------------: | :--------------: | :-------------------------------------------------------------------------------------: |
|             Prepared              |        0         |               The transaction is prepared, and is ready to be completed.                |
|              Started              |        1         |         The transaction has been started. This is used at reservation payments.         |
|             Succeeded             |        2         |                       The transaction was successfully completed.                       |
|              Timeout              |        3         |                             The transaction has timed out.                              |
|           ShopIsDeleted           |        4         |         The shop that created the transaction has been deleted in the meantime.         |
|           ShopIsClosed            |        5         |         The shop that created the transaction has been closed in the meantime.          |
|             Rejected              |        6         |                           The user rejected the transaction.                            |
|          RejectedByShop           |        12        |                       The transaction was cancelled by the shop.                        |
|              Storno               |        13        |                        Storno amount for a previous transaction.                        |
|             Reserved              |        14        |                        The transaction amount has been reserved.                        |
|              Deleted              |        15        |                              The transaction was deleted.                               |
|              Expired              |        16        |                              The transaction has expired.                               |
|            Authorized             |        17        |            The card payment transaction is authorized but not captured yet.             |
|             Reversed              |        18        | The authorization was reversed (either captured with amount 0, or net captured at all). |
|       InvalidPaymentRecord        |       210        |                   A payment to the given transaction does not exists.                   |
|          PaymentTimeOut           |       211        |                      The payment of the transaction has timed out.                      |
|       InvalidPaymentStatus        |       212        |                 The payment of the transaction is in an invalid status.                 |
| PaymentSenderOrRecipientIsInvalid |       213        |            The sender or recipient user was not found in the Barion system.             |
|              Unknown              |       255        |                         The transaction is in an unknown state.                         |
