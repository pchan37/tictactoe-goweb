<template>
  <svg :width="width" :viewBox="viewBox" xmlns="http://www.w3.org/2000/svg" id="game">
    <line x1="33%" y1="100%" x2="33%" y2="0%" stroke="black"></line>
    <line x1="67%" y1="100%" x2="67%" y2="0%" stroke="black"></line>

    <line x1="0%" y1="33%" x2="100%" y2="33%" stroke="black"></line>
    <line x1="0%" y1="67%" x2="100%" y2="67%" stroke="black"></line>
  </svg>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';
import SVGHelpers from '@/SVGLibrary';

export default {
  name: 'tic-tac-toe',
  mixins: [SVGHelpers],
  mounted() {
    const svgGameElem = document.getElementById('game');
    svgGameElem.addEventListener('click', this.playerTurn, true);
    this.setMessage('I am thinking...');
    setTimeout(this.nextTurn, 100);
  },
  computed: {
    ...mapGetters('game', [
      'OPPONENT_STAGE',

      'board',
      'emptyToken',
      'opponent',
      'tokens',
      'turn',
    ]),
  },
  methods: {
    ...mapMutations('game', [
      'setGameStage',
      'setMessage',
      'setTurn',
    ]),
    nextTurn() {
      this.setMessage(`It is ${this.tokens[this.turn]}'s turn`);
      if (this.opponent.isAI && this.opponent.token === this.tokens[this.turn]) {
        const svgGameElem = document.getElementById('game');
        svgGameElem.removeEventListener('click', this.playerTurn, true);
        this.computerTurn();
        svgGameElem.addEventListener('click', this.playerTurn, true);
      }
    },
    playerTurn(evt) {
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
        this.setMessage('Sorry, that is not a valid play!  Try again!');
      }
    },
    computerTurn() {
      const [, bestMove] = this.minimax(this.board, this.turn, this.turn, 0);

      this.drawToken(bestMove[1], bestMove[0], this.tokens[this.turn]);
      this.board[bestMove[0]][bestMove[1]] = this.tokens[this.turn];
      this.endTurn();
    },
    minimax(board, initialTurn, turn, depth) {
      if (this.tied()) {
        return [0 - depth, null];
      }
      if (this.someoneWon()) {
        return (turn !== initialTurn) ? [Number.MAX_SAFE_INTEGER - depth, null]
          : [Number.MIN_SAFE_INTEGER + depth, null];
      }

      const moves = [];
      const scores = [];
      for (let i = 0; i < 3; i += 1) {
        for (let j = 0; j < 3; j += 1) {
          if (board[i][j] === this.emptyToken) {
            moves.push([i, j]);
          }
        }
      }

      moves.forEach((move) => {
        board[move[0]][move[1]] = this.tokens[turn];
        const [score] = this.minimax(board, initialTurn, 1 - turn, depth + 1);
        scores.push(score);
        board[move[0]][move[1]] = this.emptyToken;
      });

      if (depth === 0) {
        this.shuffle(scores, moves);
      }
      if (turn === initialTurn) {
        const bestScoreIndex = scores.indexOf(Math.max(...scores));
        return [scores[bestScoreIndex], moves[bestScoreIndex]];
      }
      const worstScoreIndex = scores.indexOf(Math.min(...scores));
      return [scores[worstScoreIndex], moves[worstScoreIndex]];
    },
    endTurn() {
      if (this.someoneWon()) {
        const svgGameElem = document.getElementById('game');
        svgGameElem.removeEventListener('click', this.playerTurn, true);

        this.setMessage(`Game Over: ${this.tokens[this.turn]} has won!`);
        this.drawGameOverScreen();
        return;
      }
      if (this.tied()) {
        const svgGameElem = document.getElementById('game');
        svgGameElem.removeEventListener('click', this.playerTurn, true);

        this.setMessage('The game has ended in a tie');
        this.drawGameOverScreen();
        return;
      }

      this.setTurn(1 - this.turn);
      this.nextTurn();
    },
    drawToken(xIndex, yIndex, token) {
      const svgGameElem = document.getElementById('game');
      if (token === 'X') {
        const leftDiag = this.createSVGElement('line', {
          x1: `${xIndex * 33 + 10}`,
          y1: `${yIndex * 33 + 10}`,
          x2: `${xIndex * 33 + 23}`,
          y2: `${yIndex * 33 + 23}`,
          stroke: 'black',
        });
        const rightDiag = this.createSVGElement('line', {
          x1: `${xIndex * 33 + 10}`,
          y1: `${yIndex * 33 + 23}`,
          x2: `${xIndex * 33 + 23}`,
          y2: `${yIndex * 33 + 10}`,
          stroke: 'black',
        });

        svgGameElem.append(leftDiag);
        svgGameElem.append(rightDiag);
      } else if (token === 'O') {
        const circle = this.createSVGElement('circle', {
          cx: `${xIndex * 33 + 0.5 * 33}`,
          cy: `${yIndex * 33 + 0.5 * 33}`,
          r: '8',
          stroke: 'black',
          fill: 'none',
        });
        svgGameElem.append(circle);
      } else {
        this.setMessage(`Unable to process ${token}`);
      }
    },
    drawGameOverScreen() {
      const svgGameElem = document.getElementById('game');

      const overlay = this.createSVGElement('rect', {
        width: '100%',
        height: '100%',
        fill: 'rgba(255, 255, 255, .7)',
      });
      svgGameElem.append(overlay);

      const playAgain = this.createSVGText('Play Again?', {
        x: '13%',
        y: '33%',
      });
      svgGameElem.append(playAgain);

      const yes = this.createSVGText('Yes!', {
        x: '20%',
        y: '67%',
        'font-size': '10',
      });
      yes.classList.add('menu-item');
      yes.addEventListener('click', () => {
        while (svgGameElem.children.length > 4) {
          svgGameElem.removeChild(svgGameElem.lastChild);
        }
        this.setMessage('');
        this.setGameStage(this.OPPONENT_STAGE);
      });
      svgGameElem.append(yes);

      const no = this.createSVGText('No', {
        x: '65%',
        y: '67%',
        'font-size': '10',
      });
      no.classList.add('menu-item');
      no.addEventListener('click', () => {
        this.setGameStage(this.OPPONENT_STAGE);
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
    isEqual(a, b, c) {
      return a !== this.emptyToken && a === b && b === c;
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
  props: ['width', 'viewBox'],
};
</script>

<style lang="scss" scoped>
.menu-item:hover {
  cursor: pointer;
  fill: blue;
}
svg {
  display: block;
  margin: 0 auto;
}
</style>
