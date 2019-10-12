<template>
    <div class="home">
        <div class="todo-image">
            <img alt="Vue logo" src="../assets/todo.png">
            <h3>今日土豆</h3>
            <p>明日复明日，明日何其多。——《明日歌》</p>
        </div>
        <a-row>
            <a-col :xs="1" :sm="3" :md="6" :lg="7" :xl="8"></a-col>
            <a-col :xs="22" :sm="18" :md="12" :lg="10" :xl="8">


                <a-input placeholder="添加一条todo" v-model="title" @pressEnter="addTodo"/>
                <div v-for="todo in todos">
                    <todo :title="todo.title" :id="todo.ID" :completed="todo.completed"></todo>
                </div>
            </a-col>
            <a-col :xs="1" :sm="3" :md="6" :lg="7" :xl="8"></a-col>
        </a-row>


    </div>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'
    import todo from '@/components/Todo.vue'
    import axios from 'axios'

    export default {
        name: 'home',
        components: {
            HelloWorld,
            todo
        },
        data() {
            return {
                title: "",
                total: 0,
                todos: "",
            }
        },
        inject: ['reload'],
        methods: {
            getAllTodo() {
                axios.get('/api/v1/todos/')
                    .then((res) => {
                        this.total = res.data.total;
                        this.todos = res.data.data.todos;
                        console.info(this.todos);
                    }).catch((error) => {
                    console.error(error)
                });
            },
            getTodo() {
            },
            addTodo() {
                let data = new FormData();
                data.append('title', this.title);
                axios.post('/api/v1/todos/', data)
                    .then((res) => {
                        console.info(res);
                        this.reload();
                        this.$notify({
                            group: 'foo',
                            type: 'success',
                            text: '添加成功'
                        });
                    }).catch((error) => {
                    console.error(error)
                });
            },

            updateTodo() {
            },
        },
        mounted() {
            this.getAllTodo()
        },

    }
</script>

<style lang="less" scoped>


</style>
