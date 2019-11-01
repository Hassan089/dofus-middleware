let cells = [];

line = 0;
row = 0;

for (let i = 15; i < 500; i++) {
    cells = [...cells, {
        id: i,
        line: line,
        linePair: line % 2 == 0,
        row: row,
        fighter: null
    }];

    if (line % 2) {
        if (row == 14) {
            line++;
            row = 0;
        } else {
            row++;
        }
    } else {
        if (row == 13) {
            line++;
            row = 0;
        } else {
            row++;
        }
    }
}

console.log(cells);

var Data = {
    characters: [],
    options: {},
    message: 'Hello',
    cells: [...cells],
};

var app = new Vue({
    el: '#app',
    data: Data,
    methods: {
        focusCharacter: function(name) {
            console.log('focus', name);
            ws.send('FOCUS|' + name)
        },
        setOption: function(name, value) {
            console.log('setOption', name, value);
            ws.send('SET_OPTION|' + name + "|" + value)
        },
        setCharacterOption: function(optionName, characterName, value) {
            console.log('setOption', optionName, characterName, value);
            ws.send('SET_CHARACTER_OPTION|' + characterName + "|" + optionName + "|" + value)
        },
    }
})

const {ipcRenderer} = require('electron');

ipcRenderer.on('focus-character', (event, id) => {
    console.log('ipc message focus character', id);
    let name = '';
    switch (id) {
        case '1': name = 'Doelia'; break;
        case '2': name = 'Lotahi'; break;
        case '3': name = 'Hystelia'; break;
        case '4': name = 'Faycott'; break;
        case '5': name = 'Truecandle'; break;
        case '6': name = 'Nornut'; break;
    }
    console.log('keyboard focus', name);
    ws.send('FOCUS|' + name)
});


