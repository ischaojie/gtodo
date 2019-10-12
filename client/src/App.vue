<template>
    <div id="app">
        <notifications group="foo"
                       position="top center"
                       classes="noti"
                       :max="3"
                       :width="400"/>
        <div id="nav">
            <router-link to="/">Home</router-link>
            |
            <router-link to="/health">Health</router-link>
            |
            <router-link to="/about">About</router-link>
        </div>

        <router-view v-if="isRouterAlive"/>
    </div>
</template>

<style lang="less">
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    #nav {
        padding: 30px;

        a {
            font-weight: bold;
            color: #2c3e50;

            &.router-link-exact-active {
                /*color: #42b983;*/
            }
        }
    }

    .noti {
        margin: 10px;
        margin-bottom: 0;
        border-radius: 3px;
        font-size: 13px;
        padding: 10px 20px;
        color: #495061;
        background: #EAF4FE;
        border: 1px solid #D4E8FD;
        .notification-title {
            letter-spacing: 1px;
            text-transform: uppercase;
            font-size: 10px;
            color: #2589F3;
        }

        .notification-content {
            // Style for content
        }
    }
</style>

<script>
    export default {
        name: 'App',
        provide() {
            return {
                reload: this.reload
            }
        },
        data() {
            return {
                isRouterAlive: true,
            }
        },
        methods: {
            reload() {
                this.isRouterAlive = false
                this.$nextTick(function () {
                    this.isRouterAlive = true
                })
            },
        }
    }
</script>