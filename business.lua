require "computelib"

function computefile(filename,sheetname)
    p = datarows.new(filename, sheetname)
    print(p:get("B6"))
    print(p:get("A6"))
    print(add(p:get("E6"),p:get("F6")))
    p:set("C6","helloworld")
    p:save()
end