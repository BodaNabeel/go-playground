module example.com/main

go 1.22.0

replace example.com/intro => ../intro

replace example.com/add => ../add

replace example.com/calculation => ../calculation

require example.com/calculation v0.0.0-00010101000000-000000000000
