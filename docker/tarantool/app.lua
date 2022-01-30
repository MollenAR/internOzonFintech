box.cfg{listen = 3302}
box.schema.user.passwd('pass')

box.once('init', function()
    s = box.schema.space.create('urls')
    s:format({
        {name = 'short_url', type = 'string'},
        {name = 'original_url', type = 'string'}
    })
    s:create_index('primary', {type = 'tree', parts = {'short_url'}})
    s:create_index('secondary', {type = 'tree', unique = true, parts = {'original_url'}})
end)
