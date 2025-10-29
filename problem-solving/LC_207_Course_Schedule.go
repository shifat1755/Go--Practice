func canFinish(numCourses int, prerequisites [][]int) bool {
    adj :=map[int][]int{}
    for i:=0; i<len(prerequisites);i++{
        a, b:=prerequisites[i][0], prerequisites[i][1]
        adj[a]=append(adj[a],b)
    }
    visit:=make([]int, numCourses)
    for i:=0;i<numCourses;i++{
        if !dfs(i,adj,visit){
            return false
        }
    }
    return true
}

func dfs(node int, adj map[int][]int, visit[]int)bool{
    if visit[node]==1{
        return true
    }
    if visit[node]==2{
        return false
    }
    visit[node]=2
    for _,nei :=range adj[node]{
        if !dfs(nei,adj,visit){
            return false
        }
    }
    visit[node]=1
    return true
}