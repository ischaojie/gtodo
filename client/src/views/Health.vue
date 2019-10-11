<template>
    <div>
        <h3>cpu使用量:</h3>
        <p>status: {{cpu.status}}, info: {{cpu.info}}</p>

        <h3>ram 使用量:</h3>
        <p>status: {{ram.status}}, info: {{ram.info}}</p>

        <h3>disk 使用量:</h3>
        <p>status: {{disk.status}}, info: {{disk.info}}</p>
    </div>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "Health",
        data() {
            return {
                sd: "health",
                cpu: "",
                ram: "",
                disk: ""
            }
        },
        methods: {
            getHealth() {
                axios.all([
                    axios.get("/api/sd/cpu"),
                    axios.get("/api/sd/disk"),
                    axios.get("/api/sd/ram"),

                ]).then(axios.spread((cpu, disk, ram) => {
                    this.cpu = cpu.data;
                    this.disk = disk.data;
                    this.ram = ram.data;
                })).catch((error) => {
                    console.error(error);
                });
            }
        },


        mounted() {
            this.getHealth()
        }

    }
</script>

<style scoped>

</style>