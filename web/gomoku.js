getBoard().then((board) => {
  for (var i = 0; i < board.length; i++) {
    var row = document.createElement("div");
    row.className = "row";
    for (var j = 0; j < board[i].length; j++) {
      var cell = document.createElement("div");
      cell.className = "cell";
      cell.id = i + "-" + j;
      cell.index = [i, j];
      cell.addEventListener("click", (event) => {
        clickTile(event.target.index[0], event.target.index[1]);
      });
      row.appendChild(cell);
    }
    document.getElementById("board").appendChild(row);
  }
});

let clicker = 0

function clickTile(col, row) {
  makeMove(col, row, (clicker % 2) + 1).then((index) => {
    if (index) {
      clicker++
      document
        .getElementById(index[0][0] + "-" + index[0][1])
        .classList.add(clicker % 2 == 1 ? "black" : "white");
      checkWin().then((result) => {
        if (result != 0) {
          var x = document.createElement("h1");
          var t = document.createTextNode(result == 1 ? "black wins" : "white wins");
          x.appendChild(t);
          document.body.appendChild(x);
        }
      })
    }
  });
}
