-- Generated by github.com/davyxu/tabtoy
-- Version: 2.8.10

local tab = {
	TMusic = {
		{ Id = 1, Pos = "非战斗界面背景音乐", Name = "Music_ZhuangYuan" 	},
		{ Id = 2, Pos = "战斗界面背景音乐", Name = "Music_City" 	},
		{ Id = 3, Pos = "按钮点击", Name = "UI_Click_Default" 	},
		{ Id = 4, Pos = "选中商品音效", Name = "UI_MenuOpen" 	},
		{ Id = 5, Pos = "返回或者关闭按钮音效", Name = "UI_MenuClose" 	},
		{ Id = 6, Pos = "go和sotp按钮音效", Name = "UI_Click_Login" 	},
		{ Id = 7, Pos = "主角跳跃音效", Name = "UI_ZYPickup" 	},
		{ Id = 8, Pos = "获得金币音效", Name = "UI_StarSucceed" 	},
		{ Id = 9, Pos = "获得小奖励音效", Name = "UI_LevelUp" 	},
		{ Id = 10, Pos = "获得大奖音效", Name = "Music_Victory" 	}
	}

}


-- Id
tab.TMusicById = {}
for _, rec in pairs(tab.TMusic) do
	tab.TMusicById[rec.Id] = rec
end

tab.Enum = {
}

return tab