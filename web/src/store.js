import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const gameModule = {
  namespaced: true,
  state: {
    OPPONENT_STAGE: 'chooseOpponent',
    TOKEN_STAGE: 'chooseToken',
    GAME_STAGE: 'playGame',

    board: null,
    emptyToken: ' ',
    message: '',
    opponent: {
      isAI: false,
      token: null,
    },
    stage: 'chooseOpponent',
    tokens: ['X', 'O'],
    turn: null,
  },
  mutations: {
    resetBoard(state) {
      state.board = [
        [state.emptyToken, state.emptyToken, state.emptyToken],
        [state.emptyToken, state.emptyToken, state.emptyToken],
        [state.emptyToken, state.emptyToken, state.emptyToken],
      ];
    },
    setGameStage(state, stage) {
      state.stage = stage;
    },
    setMessage(state, message) {
      state.message = message;
    },
    setOpponentIsAI(state, isAI) {
      state.opponent.isAI = isAI;
    },
    setOpponentToken(state, token) {
      state.opponent.token = token;
    },
    setTurn(state, turn) {
      state.turn = turn;
    },
  },
  actions: {
  },
  getters: {
    OPPONENT_STAGE: state => state.OPPONENT_STAGE,
    TOKEN_STAGE: state => state.TOKEN_STAGE,
    GAME_STAGE: state => state.GAME_STAGE,

    board: state => state.board,
    emptyToken: state => state.emptyToken,
    message: state => state.message,
    opponent: state => state.opponent,
    stage: state => state.stage,
    tokens: state => state.tokens,
    turn: state => state.turn,
  },
};

export default new Vuex.Store({
  modules: {
    game: gameModule,
  },
  state: {

  },
  mutations: {

  },
  actions: {

  },
});
