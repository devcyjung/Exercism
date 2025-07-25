use std::collections::HashMap;
use graph::{Graph, Edge, Node};

pub mod graph {
    pub use graph_items::{
        edge::Edge,
        node::Node,
    };

    #[derive(Clone, Default, Debug, PartialEq, Eq)]
    pub struct Graph {
        pub nodes: Vec<Node>,
        pub edges: Vec<Edge>,
        pub attrs: crate::HashMap<String, String>,
    }

    pub mod graph_items {
        pub mod edge {
            #[derive(Clone, Default, Debug, PartialEq, Eq)]
            pub struct Edge {
                pub src: String,
                pub dst: String,
                pub attrs: crate::HashMap<String, String>,
            }
        }

        pub mod node {
            #[derive(Clone, Default, Debug, PartialEq, Eq)]
            pub struct Node {
                pub name: String,
                pub attrs: crate::HashMap<String, String>,
            }
        }
    }
}

impl Graph {
    pub fn new() -> Self {
        Self::default()
    }

    pub fn with_edges(mut self, edges: &[Edge]) -> Self {
        self.edges = edges.to_vec();
        self
    }

    pub fn with_nodes(mut self, nodes: &[Node]) -> Self {
        self.nodes = nodes.to_vec();
        self
    }

    pub fn with_attrs(mut self, attrs: &[(&str, &str)]) -> Self {
        self.attrs = attrs.iter().map(|&(k, v)| (k.into(), v.into())).collect();
        self
    }

    pub fn node(&self, name: &str) -> Option<&Node> {
        self.nodes.iter().find(|n| n.name == name)
    }
}

impl Edge {
    pub fn new(src: &str, dst: &str) -> Self {
        Self {
            src: src.into(),
            dst: dst.into(),
            ..Self::default()
        }
    }

    pub fn with_attrs(mut self, attrs: &[(&str, &str)]) -> Self {
        self.attrs = attrs.iter().map(|&(k, v)| (k.into(), v.into())).collect();
        self
    }

    pub fn attr(&self, name: &str) -> Option<&str> {
        self.attrs.get(name).map(String::as_str)
    }
}

impl Node {
    pub fn new(name: &str) -> Self {
        Self {
            name: name.into(),
            ..Self::default()
        }
    }

    pub fn with_attrs(mut self, attrs: &[(&str, &str)]) -> Self {
        self.attrs = attrs.iter().map(|&(k, v)| (k.into(), v.into())).collect();
        self
    }

    pub fn attr(&self, name: &str) -> Option<&str> {
        self.attrs.get(name).map(String::as_str)
    }
}