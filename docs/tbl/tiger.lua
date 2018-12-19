// Generated by github.com/davyxu/tabtoy
// Version: 2.8.10

module table {
export var TTiger : table.ITTigerDefine[] = [
		{ Id : 11, Number : 1, Weight : 10 	},
		{ Id : 12, Number : 2, Weight : 3 	},
		{ Id : 13, Number : 3, Weight : 1 	},
		{ Id : 14, Number : 0, Weight : 7 	},
		{ Id : 21, Number : 1, Weight : 10 	},
		{ Id : 22, Number : 2, Weight : 2 	},
		{ Id : 23, Number : 3, Weight : 2 	},
		{ Id : 24, Number : 0, Weight : 6 	},
		{ Id : 31, Number : 1, Weight : 10 	},
		{ Id : 32, Number : 2, Weight : 3 	},
		{ Id : 33, Number : 3, Weight : 1 	},
		{ Id : 34, Number : 0, Weight : 5 	}
	]


// Id
export var TTigerById : game.Dictionary<table.ITTigerDefine> = {}
function readTTigerById(){
  for(let rec of TTiger) {
    TTigerById[rec.Id] = rec; 
  }
}
readTTigerById();
}

