<template>
  <v-container fluid>
    <v-layout row wrap align-center justify-space-between>
      <v-flex xs12>
        <svg :width="svgWidth" viewBox="0 0 99 99" xmlns="http://www.w3.org/2000/svg"
             v-if="stage === 'chooseOpponent'">
          <text x="20%" y="33%" fill="blue" font-size="10">Play against?</text>
          <text x="15%" y="67%" fill="black" font-size="10" class="menu-item"
                @click="opponentChosen(human);">
            Person
          </text>
          <text x="70%" y="67%" fill="black" font-size="10" class="menu-item"
                @click="opponentChosen(ai);">
            AI
          </text>
        </svg>
        <svg :width="svgWidth" viewBox="0 0 99 99" xmlns="http://www.w3.org/2000/svg"
             v-else-if="stage === 'chooseToken'">
          <text x="5" y="33%" fill="blue" font-size="10">Choose your token</text>
          <text x="30%" y="67%" fill="black" font-size="10" class="menu-item"
                @click="tokenChosen('X');">
            X
          </text>
          <text x="65%" y="67%" fill="black" font-size="10" class="menu-item"
                @click="tokenChosen('O');">
            O
          </text>
        </svg>
        <svg :width="svgWidth" viewBox="0 0 99 99" xmlns="http://www.w3.org/2000/svg"
             v-else-if="stage === 'playGame'" id="game">
          <line x1="33%" y1="100%" x2="33%" y2="0" stroke="black"></line>
          <line x1="67%" y1="100%" x2="67%" y2="0" stroke="black"></line>

          <line x1="0" y1="33%" x2="100%" y2="33%" stroke="black"></line>
          <line x1="0" y1="67%" x2="100%" y2="67%" stroke="black"></line>
        </svg>
      </v-flex>
      <v-flex xs12 mt-4 text-xs-center>
        <span class="headline">{{ message }}</span>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: 'game',
  data() {
    return {
      ai: 'AI',
      board: null,
      emptyToken: ' ',
      human: 'Human',
      tokens: ['X', 'O'],
      turn: null,
      message: '',
      opponent: {},
      stage: 'chooseOpponent',
      windowWidth: window.innerWidth,
    };
  },
  mounted() {
    window.addEventListener('resize', () => {
      this.windowWidth = window.innerWidth;
    });
  },
  computed: {
    isMobile() {
      return this.windowWidth <= 768;
    },
    isTablet() {
      return this.windowWidth > 768 && this.windowWidth <= 960;
    },
    isLaptopOrBigger() {
      return this.windowWidth > 960;
    },
    svgWidth() {
      if (this.isMobile) {
        return '100%';
      }
      if (this.isTablet) {
        return '40%';
      }
      return '33%';
    },
  },
  methods: {
    setup() {
      this.board = [
        [this.emptyToken, this.emptyToken, this.emptyToken],
        [this.emptyToken, this.emptyToken, this.emptyToken],
        [this.emptyToken, this.emptyToken, this.emptyToken],
      ];
    },
    opponentChosen(opponent) {
      this.opponent.isAi = (opponent === this.ai);
      if (this.opponent.isAi && Math.random() < 0.5) {
        this.turn = Math.floor(Math.random() * 2);
        this.opponent.token = this.tokens[this.turn];
        this.stage = 'playGame';
        this.setup();
        this.$nextTick(() => {
          this.computerPlayToken();
          const svgGameElem = document.getElementById('game');
          svgGameElem.addEventListener('click', this.playToken, true);
        });
      } else {
        this.stage = 'chooseToken';
      }
    },
    tokenChosen(token) {
      this.stage = 'playGame';
      this.setup();
      this.turn = (token === this.tokens[0]) ? 0 : 1;
      this.opponent.token = this.tokens[1 - this.turn];

      this.message = `It is ${this.tokens[this.turn]}'s turn`;
      this.$nextTick(() => {
        const svgGameElem = document.getElementById('game');
        svgGameElem.addEventListener('click', this.playToken, true);
      });
    },
    computerPlayToken() {
      const [, bestMove] = this.minimax(this.board, this.turn, this.turn, 0);

      this.drawToken(bestMove[1], bestMove[0], this.tokens[this.turn]);
      this.board[bestMove[0]][bestMove[1]] = this.tokens[this.turn];
      this.endTurn();
    },
    minimax(board, initialTurn, turn, depth) {
      // check if game is over
      if (this.tied()) {
        return [0 - depth, null];
      }
      if (this.someoneWon()) {
        return (turn !== initialTurn) ? [Number.MAX_SAFE_INTEGER - depth, null]
          : [Number.MIN_SAFE_INTEGER + depth, null];
      }

      // get available moves
      const moves = [];
      const scores = [];
      for (let i = 0; i < 3; i += 1) {
        for (let j = 0; j < 3; j += 1) {
          if (board[i][j] === this.emptyToken) {
            moves.push([i, j]);
          }
        }
      }

      // play each available move and evaluate for a score
      moves.forEach((move) => {
        board[move[0]][move[1]] = this.tokens[turn];
        const [score] = this.minimax(board, initialTurn, 1 - turn, depth + 1);
        scores.push(score);
        board[move[0]][move[1]] = this.emptyToken;
      });
      // return the best move
      this.shuffle(scores, moves);
      if (turn === initialTurn) {
        const bestScoreIndex = scores.indexOf(Math.max(...scores));
        return [scores[bestScoreIndex], moves[bestScoreIndex]];
      }
      const worstScoreIndex = scores.indexOf(Math.min(...scores));
      return [scores[worstScoreIndex], moves[worstScoreIndex]];
    },
    playToken(evt) {
      const svgGameElem = document.getElementById('game');
      const point = svgGameElem.createSVGPoint();

      point.x = evt.clientX;
      point.y = evt.clientY;

      const svgCoordinates = point.matrixTransform(svgGameElem.getScreenCTM().inverse());

      const xIndex = Math.floor(svgCoordinates.x / 33);
      const yIndex = Math.floor(svgCoordinates.y / 33);

      if (this.board[yIndex][xIndex] === this.emptyToken) {
        this.drawToken(xIndex, yIndex, this.tokens[this.turn]);
        this.board[yIndex][xIndex] = this.tokens[this.turn];
        this.endTurn();
      } else {
        this.message = 'Sorry, that is not a valid play!  Try again!';
      }
    },
    drawToken(xIndex, yIndex, token) {
      const svgGameElem = document.getElementById('game');
      if (token === 'X') {
        const leftDiag = document.createElementNS('http://www.w3.org/2000/svg', 'line');
        leftDiag.setAttributeNS(null, 'x1', `${xIndex * 33 + 10}`);
        leftDiag.setAttributeNS(null, 'y1', `${yIndex * 33 + 10}`);
        leftDiag.setAttributeNS(null, 'x2', `${xIndex * 33 + 23}`);
        leftDiag.setAttributeNS(null, 'y2', `${yIndex * 33 + 23}`);
        leftDiag.setAttributeNS(null, 'stroke', 'black');
        const rightDiag = document.createElementNS('http://www.w3.org/2000/svg', 'line');
        rightDiag.setAttributeNS(null, 'x1', `${xIndex * 33 + 10}`);
        rightDiag.setAttributeNS(null, 'y1', `${yIndex * 33 + 23}`);
        rightDiag.setAttributeNS(null, 'x2', `${xIndex * 33 + 23}`);
        rightDiag.setAttributeNS(null, 'y2', `${yIndex * 33 + 10}`);
        rightDiag.setAttributeNS(null, 'stroke', 'black');
        svgGameElem.append(leftDiag);
        svgGameElem.append(rightDiag);
      } else if (token === 'O') {
        const circle = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
        circle.setAttributeNS(null, 'cx', `${xIndex * 33 + 0.5 * 33}`);
        circle.setAttributeNS(null, 'cy', `${yIndex * 33 + 0.5 * 33}`);
        circle.setAttributeNS(null, 'r', '8');
        circle.setAttributeNS(null, 'stroke', 'black');
        circle.setAttributeNS(null, 'fill', 'none');
        svgGameElem.append(circle);
      } else {
        this.message = `Unable to process ${token}`;
      }
    },
    endTurn() {
      if (this.someoneWon()) {
        const svgGameElem = document.getElementById('game');
        svgGameElem.removeEventListener('click', this.playToken, true);

        this.message = `Game Over: ${this.tokens[this.turn]} has won!`;
        this.drawGameOverScreen();
        return;
      }
      if (this.tied()) {
        const svgGameElem = document.getElementById('game');
        svgGameElem.removeEventListener('click', this.playToken, true);

        this.message = 'The game has ended in a tie';
        this.drawGameOverScreen();
        return;
      }
      this.turn = (1 - this.turn);
      this.message = `It is ${this.tokens[this.turn]}'s turn`;
      if (this.tokens[this.turn] === this.opponent.token && this.opponent.isAi) {
        this.computerPlayToken();
      }
    },
    drawGameOverScreen() {
      const svgGameElem = document.getElementById('game');

      const overlay = document.createElementNS('http://www.w3.org/2000/svg', 'rect');
      overlay.setAttributeNS(null, 'width', '100%');
      overlay.setAttributeNS(null, 'height', '100%');
      overlay.setAttributeNS(null, 'fill', 'rgba(255, 255, 255, .7)');
      svgGameElem.append(overlay);

      const playAgainText = document.createTextNode('Play Again?');
      const playAgain = document.createElementNS('http://www.w3.org/2000/svg', 'text');
      playAgain.setAttributeNS(null, 'x', '13%');
      playAgain.setAttributeNS(null, 'y', '33%');
      playAgain.appendChild(playAgainText);
      svgGameElem.append(playAgain);

      const yesText = document.createTextNode('Yes!');
      const yes = document.createElementNS('http://www.w3.org/2000/svg', 'text');
      yes.setAttributeNS(null, 'x', '20%');
      yes.setAttributeNS(null, 'y', '67%');
      yes.setAttributeNS(null, 'font-size', '10');
      yes.appendChild(yesText);
      yes.classList.add('menu-item');
      yes.addEventListener('click', () => {
        const svgGameElement = document.getElementById('game');
        while (svgGameElement.children.length > 4) {
          svgGameElement.removeChild(svgGameElement.lastChild);
        }
        this.message = '';
        this.stage = 'chooseOpponent';
      });
      svgGameElem.append(yes);

      const noText = document.createTextNode('No');
      const no = document.createElementNS('http://www.w3.org/2000/svg', 'text');
      no.setAttributeNS(null, 'x', '65%');
      no.setAttributeNS(null, 'y', '67%');
      no.setAttributeNS(null, 'font-size', '10');
      no.appendChild(noText);
      no.classList.add('menu-item');
      no.addEventListener('click', () => {
        this.$router.push('/');
      });
      svgGameElem.append(no);
    },
    someoneWon() {
      return (this.isEqual(this.board[0][0], this.board[0][1], this.board[0][2])
              || this.isEqual(this.board[1][0], this.board[1][1], this.board[1][2])
              || this.isEqual(this.board[2][0], this.board[2][1], this.board[2][2])
              || this.isEqual(this.board[0][0], this.board[1][0], this.board[2][0])
              || this.isEqual(this.board[0][1], this.board[1][1], this.board[2][1])
              || this.isEqual(this.board[0][2], this.board[1][2], this.board[2][2])
              || this.isEqual(this.board[0][0], this.board[1][1], this.board[2][2])
              || this.isEqual(this.board[0][2], this.board[1][1], this.board[2][0]));
    },
    isEqual(a, b, c) {
      return a !== this.emptyToken && a === b && b === c;
    },
    tied() {
      let boardFilled = true;
      for (let i = 0; i < 3; i += 1) {
        for (let j = 0; j < 3; j += 1) {
          if (this.board[i][j] === this.emptyToken) {
            boardFilled = false;
          }
        }
      }
      return boardFilled && !this.someoneWon();
    },
    shuffle(arrayOne, arrayTwo) {
      let index = arrayOne.length;
      let randomIndex;
      let tempOne;
      let tempTwo;

      while (index) {
        randomIndex = Math.floor(Math.random() * index);
        index -= 1;
        tempOne = arrayOne[index];
        tempTwo = arrayTwo[index];
        arrayOne[index] = arrayOne[randomIndex];
        arrayTwo[index] = arrayTwo[randomIndex];
        arrayOne[randomIndex] = tempOne;
        arrayTwo[randomIndex] = tempTwo;
      }
    },
  },
};
</script>

<style lang="scss">
.menu-item:hover {
  cursor: pointer;
  fill: blue;
}
svg {
  display: block;
  margin: 0 auto;
}
</style>
