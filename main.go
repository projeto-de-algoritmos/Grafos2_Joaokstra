package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    var option string
	var g Graph
    fillGraph(&g)
    fmt.Println("Grafo carregado com sucesso")
    g.String()
    for{
        fmt.Println("\n ------------------")
        fmt.Println("Selecione uma opção: ")
        fmt.Print(" 1 - Inserir cidade \n 2 - Inserir conexão entre cidades \n 3 - Imprimir lista de adjacência \n 4 - Traçar menor caminho entre cidades \n")
        scanner.Scan()
        option = scanner.Text()
        switch option{
        case "1":
            fmt.Print("Insira o nome da cidade: ")
            scanner.Scan()
            name := scanner.Text()
            node := Node{Key: name}
            g.AddNode(&node)
        case "2":
            fmt.Print("Insira o nome da cidade: ")
            scanner.Scan()
            n1 := g.GetNode(scanner.Text())
            fmt.Print("Insira o nome da cidade vizinha: ")
            scanner.Scan()
            n2 := g.GetNode(scanner.Text())
            fmt.Print("Insira a distância: ")
            scanner.Scan()
            weight, err := strconv.Atoi(scanner.Text())
            if(err != nil){
                g.AddEdge(n1, n2, weight)
            }
        case "3":
            g.String()
        case "4":
            fmt.Print("Insira o ponto de partida: ")
            scanner.Scan()
            n1 := g.GetNode(scanner.Text())
            fmt.Print("Insira o ponto de chegada: ")
            scanner.Scan()
            n2 := g.GetNode(scanner.Text())
            if(n1 == nil || n2 == nil){
                fmt.Println("Cidade inválida")
            }else{
                fmt.Println("\n ------------------\n Realizando busca com Dijkstra")
                weight, path := g.Dijkstra(n1, n2)
                fmt.Println("Distância: " + strconv.Itoa(weight) + "\n Menor Caminho: ")
                fmt.Print(path)
            }
        }
    }
}

func printPath(path []*Node){
    str := path[len(path) -1].Key
    for i:= len(path) -2; i>=0; i--{
        str += "->" + path[i].Key
    }
    fmt.Println(str)
}

func fillGraph(g *Graph){
	nA := Node{Key: "Luziania", Value: 1}
    nB := Node{Key: "Gama", Value: 2}
    nC := Node{Key: "Valparaíso", Value: 3}
    nD := Node{Key: "Santa Maria", Value: 4}
    nE := Node{Key: "Asa Norte", Value: 5}
    nF := Node{Key: "Asa sul", Value: 6}
    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    g.AddNode(&nD)
    g.AddNode(&nE)
    g.AddNode(&nF)

    g.AddEdge(&nA, &nC, 20)
    g.AddEdge(&nB, &nC, 20)
    g.AddEdge(&nB, &nD, 5)
    g.AddEdge(&nC, &nD, 15)
    g.AddEdge(&nE, &nF, 10)
    g.AddEdge(&nF, &nB, 30)
}
