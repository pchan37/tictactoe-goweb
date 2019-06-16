<template>
<v-app>
  <v-navigation-drawer app mobile-break-point=102400 v-model="drawer">
    <v-list>
      <v-list-tile to="/" @click="drawer = false">
        <v-list-tile-action>
          <v-icon>fas fa-home-alt</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>TicTacToe</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile to="/game" @click="drawer = false">
        <v-list-tile-action>
          <v-icon>videogame_asset</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Play</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-divider></v-divider>
      <v-list-tile>
        <v-list-tile-action>
          <v-icon>fas fa-sign-in-alt</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Login</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
</v-navigation-drawer>

  <v-toolbar fixed clipped-left app dark color="primary">
    <v-toolbar-side-icon class="hidden-md-and-up" @click.stop="drawer = !drawer;">
    </v-toolbar-side-icon>
    <v-toolbar-title>
      <router-link to="/" class="toolbar-title">TicTacToe</router-link>
    </v-toolbar-title>

    <v-spacer></v-spacer>

    <v-toolbar-items class="hidden-sm-and-down">
      <v-btn flat to="/game">
        <v-icon>videogame_asset</v-icon>
        <span class="ml-2">Play</span>
      </v-btn>
      <v-btn flat href="/login">
        <v-icon>fas fa-sign-in-alt</v-icon>
        <span class="ml-2">Login</span>
      </v-btn>
    </v-toolbar-items>
  </v-toolbar>

  <v-content>
    <router-view></router-view>
  </v-content>
</v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
    };
  },
};
</script>


<style lang="scss">
.toolbar-title {
  color: inherit;
  text-decoration: inherit;
}
</style>
