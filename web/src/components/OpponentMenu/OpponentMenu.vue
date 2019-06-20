<template>
  <svg :width="width" :viewBox="viewBox" xmlns="http://www.w3.org/2000/svg">
    <text x="20%" y="33%" fill="black" font-size="10">Play against?</text>
    <text x="15%" y="67%" fill="black" font-size="10" class="menu-item"
          @click="opponentChosen(person);">
      {{ person }}
    </text>
    <text x="70%" y="67%" fill="black" font-size="10" class="menu-item"
          @click="opponentChosen(ai);">
      {{ ai }}
    </text>
  </svg>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';

export default {
  name: 'opponent-menu',
  mounted() {
    this.setMessage('');
  },
  computed: {
    ...mapGetters('game', [
      'TOKEN_STAGE',
      'GAME_STAGE',

      'opponent',
      'tokens',
      'turn',
    ]),
  },
  data() {
    return {
      ai: 'AI',
      person: 'Person',
    };
  },
  methods: {
    ...mapMutations('game', [
      'resetBoard',
      'setGameStage',
      'setMessage',
      'setOpponentIsAI',
      'setOpponentToken',
      'setTurn',
    ]),
    opponentChosen(opponent) {
      this.setOpponentIsAI(opponent === this.ai);
      if (this.opponent.isAI && Math.random() < 0.5) {
        this.setTurn(Math.floor(Math.random() * 2));
        this.setOpponentToken(this.tokens[this.turn]);
        this.resetBoard();
        this.setGameStage(this.GAME_STAGE);
      } else {
        this.setGameStage(this.TOKEN_STAGE);
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
