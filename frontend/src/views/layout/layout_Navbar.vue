<template>
    <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
        <b-container>
                <b-navbar-brand href="#">Gin——Blog</b-navbar-brand>

                <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

                <b-collapse id="nav-collapse" is-nav>

                <!-- Right aligned nav items -->
                <b-navbar-nav class="ml-auto">
                    <!-- <b-nav-form>
                    <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
                    <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
                    </b-nav-form> -->

                    <!-- <b-nav-item-dropdown right>
                    <template #button-content>
                        <em>User</em>
                    </template>
                    <b-dropdown-item href="#">Profile</b-dropdown-item>
                    <b-dropdown-item href="#">Sign Out</b-dropdown-item>
                    </b-nav-item-dropdown> -->
                    <b-nav-item @click="$router.replace({name:'login'})">Login</b-nav-item>
                    <b-nav-item @click="$router.replace({name:'register'})">register</b-nav-item>
                </b-navbar-nav>
            </b-collapse>
        </b-container>
    </b-navbar>
    </div>
  </template>
<script>

import { mapState, mapActions } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: mapActions('userModule', ['logout']),
};
</script>
<style scoped>
</style>
