-- Generated by github.com/davyxu/tabtoy
-- Version: 2.8.10

local tab = {
	TTurntableNew = {
		{ Id = 1, Nums = "0;1;2;3;4;5;6;1;2;0", Max = 6 	},
		{ Id = 2, Nums = "0;1;2;3;4;5;6;7;8;9", Max = 9 	}
	}

}


-- Id
tab.TTurntableNewById = {}
for _, rec in pairs(tab.TTurntableNew) do
	tab.TTurntableNewById[rec.Id] = rec
end

tab.Enum = {
}

return tab