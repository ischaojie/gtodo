<template>
    <div class="todo" v-if="completed===0">
        <a-radio @click="updateTodo"></a-radio>{{title}}
        <div style="flex: 1; text-align: right;">
            <a-icon type="delete" theme="twoTone" @click="deleteTodo"/>
        </div>
    </div>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "todo",
        props: ['id', 'title', 'completed'],
        inject: ['reload'],
        methods: {
            deleteTodo() {
                axios.delete('/api/v1/todos/' + this.id)
                    .then((res) => {
                        this.reload();
                        this.$notify({
                            group: 'foo',
                            type: 'info',
                            text: '删除成功'
                        });
                    }).catch((error) => {
                    console.error(error)
                });
            },
            updateTodo() {
                let data = new FormData();
                data.append('title', this.title);
                data.append('completed', 1);
                axios.put('/api/v1/todos/' + this.id, data)
                    .then((res) => {
                        this.reload();
                        this.$notify({
                            group: 'foo',
                            type: 'info',
                            text: '好棒耶！完成任务啦！'
                        });
                    }).catch((error) => {
                    console.error(error)
                });
            },
        }

    }
</script>

<style scoped lang="less">
    .todo {
        margin: 8px auto;
        padding: 12px 8px;
        text-align: left;
        -moz-box-shadow: 0px 0px 2px #847070;
        -webkit-box-shadow: 0px 0px 2px #847070;
        box-shadow: 0px 0px 2px #847070;
        border-radius: 4px;
        display: flex;
        align-items: center;
    }

    .completed {
        text-decoration: line-through;
    }
</style>