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

function clickTile(col, row) {
  makeMove(col, row).then((index) => {
    if (index) {
      document
        .getElementById(index[0][0] + "-" + index[0][1])
        .classList.add("black");
      document
        .getElementById(index[1][0] + "-" + index[1][1])
        .classList.add("white");
    }
  });
}
