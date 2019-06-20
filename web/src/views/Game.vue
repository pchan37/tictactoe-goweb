<template>
  <v-container fluid>
    <v-layout row wrap align-center justify-space-between>
      <v-flex xs12>
        <opponent-menu :width="svgWidth" view-box="0 0 99 99"
                       v-if="stage === OPPONENT_STAGE"></opponent-menu>
        <token-menu :width="svgWidth" view-box="0 0 99 99"
                    v-else-if="stage === TOKEN_STAGE"></token-menu>
        <tic-tac-toe :width="svgWidth" view-box="0 0 99 99"
                     v-else-if="stage === GAME_STAGE"></tic-tac-toe>
      </v-flex>
      <v-flex xs12 mt-4 text-xs-center>
        <span class="headline">{{ message }}</span>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapGetters } from 'vuex';

import OpponentMenu from '@/components/OpponentMenu/OpponentMenu.vue';
import TokenMenu from '@/components/TokenMenu/TokenMenu.vue';
import TicTacToe from '@/components/TicTacToe/TicTacToe.vue';

export default {
  name: 'game',
  data() {
    return {
      windowWidth: window.innerWidth,
    };
  },
  mounted() {
    window.addEventListener('resize', () => {
      this.windowWidth = window.innerWidth;
    });
  },
  components: {
    'opponent-menu': OpponentMenu,
    'token-menu': TokenMenu,
    'tic-tac-toe': TicTacToe,
  },
  computed: {
    ...mapGetters('game', [
      'OPPONENT_STAGE',
      'TOKEN_STAGE',
      'GAME_STAGE',

      'message',
      'stage',
    ]),
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
