<template>
<v-container fill-height>
  <v-layout row align-center justify-center>
    <v-flex xs12>
      <p class="display-1 text-xs-center">
        Sorry, I could not find the resource you requested.
        Do you want to go back to the <router-link to="/">home page</router-link>?
      </p>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
export default {
  name: 'NotFound',
};
</script>
