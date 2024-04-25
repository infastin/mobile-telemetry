local requests = {}
local lri = 0

local function random_uuid()
    local template = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
    return template:gsub("[xy]", function(c)
        local v = (c == "x") and math.random(0, 0xf) or math.random(8, 0xb)
        return string.format("%x", v)
    end)
end

local function random_string(size)
  local alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"

  local s = {}
  for _ = 1, size do
    local idx = math.random(1, #alphabet)
    s[#s+1] = alphabet:sub(idx, idx)
  end

  return table.concat(s)
end

local function random_info()
  local systems = {"ios", "android"}
  local widths = {1080, 1440, 1170, 1125, 1242}
  local heights = {1920, 2560, 2280, 3120, 2160}

  return {
    user_id = random_uuid(),
    device = {
      manufacturer = random_string(math.random(1, 16)),
      model = random_string(math.random(1, 16)),
      build_number = random_string(math.random(1, 16)),
      os = systems[math.random(1, #systems)],
      os_version = tostring(math.random(1, 20)),
      screen_width = widths[math.random(1, #widths)],
      screen_height = heights[math.random(1, #heights)],
    },
    app_version = string.gsub("x.x.x", "[xyz]", function()
      return tostring(math.random(0, 5))
    end)
  }
end

local function random_action()
  local actions = {"open", "close", "switch", "tap", "drag", "hold"}

  return {
    action = actions[math.random(1, #actions)],
    data = {
      page = math.random(1, 10),
      x = math.random(1080, 1440),
      y = math.random(1920, 3120),
    },
    timestamp = os.date("%Y-%m-%dT%H:%M:%SZ")
  }
end

local function to_json(val)
  if type(val) == "string" then
    return '"' .. val:gsub("\\", "\\\\"):gsub('"', '\\"') .. '"'
  elseif type(val) == "number" then
    return tostring(val)
  elseif type(val) == "boolean" then
    return val and 'true' or 'false'
  elseif type(val) == "nil" then
    return 'null'
  elseif type(val) == "function" or type(val) == "thread" or type(val) == "userdata" then
    return nil
  end

  local s = {}

  if val[1] == nil then
    s[#s+1] = "{"
    for k, v in pairs(val) do
      if #s > 1 then
        s[#s+1] = ","
      end
      s[#s+1] = string.format('"%s":%s', k, to_json(v))
    end
    s[#s+1] = "}"
  else
    s[#s+1] = "["
    for i, v in ipairs(val) do
      if i ~= 1 then
        s[#s+1] = ","
      end
      s[#s+1] = to_json(v)
    end
    s[#s+1] = "]"
  end

  return table.concat(s)
end

function init()
  math.randomseed(os.time())

  local method = "POST"
  local headers = {
    ["Content-Type"] = "application/json",
  }

  for _ = 1, 10000 do
    local body = {
      info = random_info(),
      data = {},
    }

    for _ = 1, math.random(3, 7) do
      body.data[#body.data+1] = random_action()
    end

    requests[#requests+1] = wrk.format(method, wrk.path, headers, to_json(body))
  end
end

function request()
  lri = (lri + 1) % (#requests + 1)
  return requests[lri]
end
