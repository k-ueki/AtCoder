package main

import "errors"

type (
	Node struct {
		name  interface{} // ノード名
		edges []*Edge     // 次に移動できるエッジ
		done  bool        // 処理済みかを表すフラグ
		cost  int         // このノードにたどり着くのに必要だったコスト
		prev  *Node       // このノードにたどりつくのに使われたノード
	}

	Edge struct {
		next *Node // 次に移動できるノード
		cost int   // 移動にかかるコスト
	}

	// 有向グラフ
	DirectedGraph struct {
		nodes map[string]*Node
	}
)

func NewNode(name interface{}) *Node {
	return &Node{
		name:  name,
		edges: []*Edge{},
		done:  false,
		cost:  -1,
		prev:  nil,
	}
}

func NewEdge(next *Node, cost int) *Edge {
	return &Edge{
		next: next,
		cost: cost,
	}
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		nodes: map[string]*Node{},
	}
}

// ノードに次の接続先を示したエッジを追加する
func (self *Node) AddEdge(edge *Edge) {
	self.edges = append(self.edges, edge)
}

// グラフの要素を追加する (接続元ノード名、接続先ノード名、移動にかかるコスト)
func (self *DirectedGraph) Add(src, dst string, cost int) {
	// ノードが既にある場合は追加しない
	srcNode, ok := self.nodes[src]
	if !ok {
		srcNode = NewNode(src)
		self.nodes[src] = srcNode
	}

	dstNode, ok := self.nodes[dst]
	if !ok {
		dstNode = NewNode(dst)
		self.nodes[dst] = dstNode
	}

	// ノードをエッジでつなぐ
	edge := NewEdge(dstNode, cost)
	srcNode.AddEdge(edge)
}

// スタートとゴールを指定して最短経路を求める
func (self *DirectedGraph) ShortestPath(start string, goal string) (ret []*Node, err error) {
	// 名前からスタート地点のノードを取得する
	startNode := self.nodes[start]

	// スタートのコストを 0 に設定することで処理対象にする
	startNode.cost = 0

	for {
		// 次の処理対象のノードを取得する
		node, err := self.nextNode()

		// 次に処理するノードが見つからなければ終了
		if err != nil {
			return nil, errors.New("Goal not found")
		}

		// ゴールまで到達した
		if node.name == goal {
			break
		}

		// 取得したノードを処理する
		self.calc(node)
	}

	// ゴールから逆順にスタートまでノードをたどっていく
	n := self.nodes[goal]
	for {
		ret = append(ret, n)
		if n.name == start {
			break
		}
		n = n.prev
	}

	return ret, nil
}

// つながっているノードのコストを計算する
func (self *DirectedGraph) calc(node *Node) {
	// ノードにつながっているエッジを取得する
	for _, edge := range node.edges {
		nextNode := edge.next

		// 既に処理済みのノードならスキップする
		if nextNode.done {
			continue
		}

		// このノードに到達するのに必要なコストを計算する
		cost := node.cost + edge.cost
		if nextNode.cost == -1 || cost < nextNode.cost {
			// 既に見つかっている経路よりもコストが小さければ処理中のノードを遷移元として記録する
			nextNode.cost = cost
			nextNode.prev = node
		}
	}

	// つながっているノードのコスト計算がおわったらこのノードは処理済みをマークする
	node.done = true
}

func (self *DirectedGraph) nextNode() (next *Node, err error) {
	// グラフに含まれるノードを線形探索する
	for _, node := range self.nodes {

		// 処理済みのノードは対象外
		if node.done {
			continue
		}

		// コストが初期値 (-1) になっているノードはまだそのノードまでの最短経路が判明していないので処理できない
		if node.cost == -1 {
			continue
		}

		// 最初に見つかったものは問答無用で次の処理対象の候補になる
		if next == nil {
			next = node
		}

		// 既に見つかったノードよりもコストの小さいものがあればそちらを先に処理しなければいけない
		if next.cost > node.cost {
			next = node
		}
	}

	// 次の処理対象となるノードが見つからなかったときはエラー
	if next == nil {
		return nil, errors.New("Untreated node not found")
	}

	return
}

func main() {

}
