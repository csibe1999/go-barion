Name information structure
This structure represents information about the name of a user in the Barion system. Fields set in the structure depend on the account type (individual or organization).

Included in
User information is used in the following structures:
UserInformation

|  Property name   | Property type |                               Description                               |
| :--------------: | :-----------: | :---------------------------------------------------------------------: |
|    LoginName     |    string     | The login name of the user. At the moment this is their e-mail address. |
|    FirstName     |    string     |               The first name of the user, if applicable.                |
|     LastName     |    string     |                The last name of the user, if applicable.                |
| OrganizationName |    string     |            The organization name of the user, if applicable.            |
