<template>
  <svg :width="width" :viewBox="viewBox" xmlns="http://www.w3.org/2000/svg">
    <text x="5" y="33%" fill="black" font-size="10">
      Choose your token
    </text>
    <text x="30%" y="67%" fill="black" font-size="10" class="menu-item"
          @click="tokenChosen(tokens[0]);">
      {{ tokens[0] }}
    </text>
    <text x="65%" y="67%" fill="black" font-size="10" class="menu-item"
          @click="tokenChosen(tokens[1]);">
      {{ tokens[1] }}
    </text>
  </svg>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';

export default {
  name: 'token-menu',
  computed: {
    ...mapGetters('game', [
      'GAME_STAGE',

      'tokens',
      'turn',
    ]),
  },
  methods: {
    ...mapMutations('game', [
      'resetBoard',
      'setGameStage',
      'setMessage',
      'setOpponentToken',
      'setTurn',
    ]),
    tokenChosen(token) {
      this.setTurn((token === this.tokens[0]) ? 0 : 1);
      this.setOpponentToken(this.tokens[1 - this.turn]);
      this.resetBoard();
      this.setMessage(`It is ${this.tokens[this.turn]}'s turn`);
      this.setGameStage(this.GAME_STAGE);
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
