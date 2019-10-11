<template>
    <div class="home">
        <img alt="Vue logo" src="../assets/todo.png">
        <a-row>
            <a-col :span="8"></a-col>
            <a-col :span="8">
                <a-input placeholder="添加一条todo" v-model="title" @pressEnter="addTodo"/>
                <div v-for="todo in todos">
                    <todo :title="todo.title" :id="todo.ID" :completed="todo.completed"></todo>
                </div>
            </a-col>
            <a-col :span="8"></a-col>
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
                todos: ""
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
                console.info(this.title)
                axios.post('/api/v1/todos/', {
                    title: this.title,
                    completed: 0
                })
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
        }
    }
</script>
