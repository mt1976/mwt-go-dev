----------------------------------------------------------------
Automatically generated  "/design/catalog/api.nfo"
----------------------------------------------------------------
Package              : n/a
Object 			     : API (api)
Endpoint 	         : API (ID)
For Project          : github.com/mt1976/mwt-go-dev/
----------------------------------------------------------------
Template Generator   : delinquentDysprosium [r4-21.12.31]
Date & Time		     : 12/12/2021 at 16:13:07
Who & Where		     : matttownsend on silicon.local
----------------------------------------------------------------
Database Information
----------------------------------------------------------------
SQL Table Name       : API
SQL Table Search Col : ID
----------------------------------------------------------------
Fields Generated
----------------------------------------------------------------
* Md  = Mandatory
* Cr  = Core Fields
* Ex  = Extra Fields
* Ov  = Override of a Core Field
* Lkp = Lookup Field
*-------------------------------------+------------+------------+----+----+----+----+----+-----------------+--------------------------+--------------------------*
| Field Name                          | Type       | Default    | Md | Cr | Ex | Ov | Lkp| Lkp Object      | Lkp Field Source         | Lkp Value                |  
*-------------------------------------+------------+------------+----+----+----+----+----+-----------------+--------------------------+--------------------------*
| ID                                  | String     |            | Y  | Y  |    |    |    |                 |                          |                          |
| Endpoint                            | String     |            | Y  | Y  |    |    |    |                 |                          |                          |
| Descr                               | String     |            | Y  | Y  |    |    |    |                 |                          |                          |
| Query                               | String     |            | Y  | Y  |    |    |    |                 |                          |                          |
| Source                              | String     |            | Y  | Y  |    |    |    |                 |                          |                          |
*-------------------------------------+------------+------------+----+----+----+----+----+-----------------+--------------------------+--------------------------*
----------------------------------------------------------------
Artifacts Generated
----------------------------------------------------------------
* application (/application/aPI_core.go)
* api (/application/aPI_api.go)
* dao (/dao/aPI_core.go)
* datamodel (/datamodel/aPI_core.go)
* catalog (/design/catalog/aPI.nfo)
----------------------------------------------------------------
End of File
----------------------------------------------------------------