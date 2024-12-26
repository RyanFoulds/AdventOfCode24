package network

import (
	"sort"
	"strings"
)

type link struct {
	first, second string
}

type network struct {
	output      string
	connections map[string]struct{}
}

type computer struct {
	name  string
	links map[string]struct{}
}

var networks map[string]struct{} = make(map[string]struct{})

var comps map[string]computer = make(map[string]computer)

var simpleComps map[string][]string = make(map[string][]string)

func ProcessLinks(linkStrs []string) {
	for _, linkStr := range linkStrs {
		computers := strings.Split(linkStr, "-")
		l := link{computers[0], computers[1]}

		//First
		fComp, okF := comps[l.first]
		if !okF {
			fComp = computer{l.first, make(map[string]struct{})}
		}
		fComp.links[l.second] = struct{}{}
		comps[fComp.name] = fComp

		//second
		sComp, okS := comps[l.second]
		if !okS {
			sComp = computer{l.second, make(map[string]struct{})}
		}
		sComp.links[l.first] = struct{}{}
		comps[sComp.name] = sComp
	}

	for _, v := range comps {
		l := make([]string, 0)
		for li := range v.links {
			l = append(l, li)
		}
		simpleComps[v.name] = l
	}
}

func FindNetworks() {
	for name, com := range simpleComps {
		linkCount := len(com)
		for i := 0; i < linkCount-1; i++ {
			for j := i + 1; j < linkCount; j++ {
				iName, jName := com[i], com[j]
				iCom := comps[iName]
				_, iContainsJ := iCom.links[jName]
				if iContainsJ {
					names := []string{name, iName, jName}
					sort.Strings(names)
					networks[strings.Join(names, ",")] = struct{}{}
				}
			}
		}
	}
}

func CountNetworks() (count int) {
	for n := range networks {
		if n[0] == 't' || n[3] == 't' || n[6] == 't' {
			count++
		}
	}
	return
}

func FindBiggestNetwork() string {
	networkCache := map[string]struct{}{}
	var retval string
	for _, v := range comps {
		longestFound := bfs(v, networkCache)
		if len(longestFound) > len(retval) {
			retval = longestFound
		}
	}

	return retval
}

func bfs(c computer, networkCache map[string]struct{}) string {
	start := network{c.name, map[string]struct{}{c.name: {}}}
	networkCache[start.output] = struct{}{}
	queue := []network{start}

	var n network
	for len(queue) > 0 {
		n, queue = queue[0], queue[1:]
		for candidate := range c.links {
			if _, ok := n.connections[candidate]; ok {
				continue
			}
			connected := true
			for existing := range n.connections {
				if _, ok := comps[existing].links[candidate]; !ok {
					connected = false
					break
				}
			}
			if connected {
				next := n.add(candidate)
				if _, ok := networkCache[next.output]; !ok {
					queue = append(queue, next)
					networkCache[next.output] = struct{}{}
				}
			}
		}
	}

	return n.output
}

func (n network) add(c string) network {
	newConnections := n.connections
	newConnections[c] = struct{}{}
	newOutput := make([]string, len(newConnections))
	i := 0
	for k := range newConnections {
		newOutput[i] = k
		i++
	}
	sort.Strings(newOutput)
	return network{strings.Join(newOutput, ","), newConnections}
}
