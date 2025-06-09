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
        Default::default()
    }

    pub fn with_edges(self, edges: &[Edge]) -> Self {
        Self {
            edges: edges.to_vec(),
            ..self
        }
    }

    pub fn with_nodes(self, nodes: &[Node]) -> Self {
        Self {
            nodes: nodes.to_vec(),
            ..self
        }
    }

    pub fn with_attrs(self, attrs: &[(&str, &str)]) -> Self {
        Self {
            attrs: attrs.iter().map(|(k, v)| ((*k).into(), (*v).into())).collect(),
            ..self
        }
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
            ..Default::default()
        }
    }

    pub fn with_attrs(self, attrs: &[(&str, &str)]) -> Self {
        Self {
            attrs: attrs.iter().map(|(k, v)| ((*k).into(), (*v).into())).collect(),
            ..self
        }
    }

    pub fn attr(&self, name: &str) -> Option<&str> {
        self.attrs.get(name).map(String::as_str)
    }
}

impl Node {
    pub fn new(name: &str) -> Self {
        Self {
            name: name.into(),
            ..Default::default()
        }
    }

    pub fn with_attrs(self, attrs: &[(&str, &str)]) -> Self {
        Self {
            attrs: attrs.iter().map(|(k, v)| ((*k).into(), (*v).into())).collect(),
            ..self
        }
    }

    pub fn attr(&self, name: &str) -> Option<&str> {
        self.attrs.get(name).map(String::as_str)
    }
}