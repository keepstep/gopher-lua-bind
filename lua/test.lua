require "function"
local http = require("http")
local json = require("json")
local client = http.client()

local function get()
    -- GET
    local request = http.request("GET", "http://cn.bing.com")
    local result, err = client:do_request(request)
    if err then
        error(err)
    end
    if not (result.code == 200) then
        error("code")
    end

    printInfo(string.len(result.body),string.sub(result.body,1,20))
    result.body = string.sub(result.body,1,32)
    printInfo(debug.traceback("", 1))
    dumptag("HTTP",result)

    local t,errD = json.decode('{"name":"god","age":10,"tag":["G","O","D"]}')
    if errD then
        printError(errD)
    else
        dumptag("JSON",t)    
    end
    t.name="human"
    t.tag[1]="H"
    table.insert(t.tag,'B')
    table.insert(t.tag,'Z')
    local s,errE = json.encode(t)
    if errE then
        printError(errE)
    else
        printInfo("JSON",s)
    end
end

-- for i = 1, 1, 1 do
--     get()
--     printInfo("----sleep----"..i)
--     sleep(2)
-- end
get()
return "ABC",1000