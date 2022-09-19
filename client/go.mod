module client

go 1.19

require (
    tmclient v0.0.0
)

replace (
    tmclient v0.0.0 => ./../tmclient/
)