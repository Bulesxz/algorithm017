class Trie(object):
    def __init__(self):
        self.root = {}
        self.end_world = '#'
    def insert(self,world):
        node = self.root
        for char in world:
            if node.has_key(char):
                node = node[char]
            else:
                node[char] = {}
        node[self.end_world]  = True
    
    def search(self,world):
        node = self.root
        for char in world:
            node = node.get(char,None)
            if not node:
                return None
        return node.end_world in node
    
    def startsWith(self, prefix):
        """
        Returns if there is any word in the trie that starts with the given prefix.
        :type prefix: str
        :rtype: bool
        """
        node = self.root
        for char in prefix:
            node = node.get(char,None)
            if not node:
                return False
        return True
        

if __name__ == '__main__':
    print ("=====================")
    trie = Trie()
    trie.insert('a')
    print trie.self('a')
