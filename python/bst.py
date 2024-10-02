from copy import deepcopy


class Node:
    def __init__(self, item: int) -> None:
        self.item: int = item
        self.left_child: Node | None = None
        self.right_child: Node | None = None


class BST:
    def __init__(self) -> None:
        self.root = None

    def insert(self, item: int) -> None:
        new_node = Node(item)

        if not self.root:
            self.root = new_node

        else:
            self._insert_child(new_node, self.root)

    def _insert_child(self, item: Node, curr: Node) -> None:
        if curr.item > item.item:
            if curr.left_child:
                self._insert_child(item, curr.left_child)
            else:
                curr.left_child = item

        else:
            if curr.right_child:
                self._insert_child(item, curr.right_child)
            else:
                curr.right_child = item

    def search(self, item: int) -> None:
        if not self.root:
            print("Binary search tree is empty!")

        else:
            found = self._search_item(item, self.root)
            print(f"{item} is {'NOT' if not found else ""} found in the tree!")

    def _search_item(self, item: int, curr: Node | None) -> bool:
        if not curr:
            return False

        if curr.item == item:
            return True

        if curr.left_child and curr.item > item:
            return self._search_item(item, curr.left_child)
        elif curr.right_child and curr.item < item:
            return self._search_item(item, curr.right_child)
        else:
            return False

    def delete(self, item: int) -> None:
        if not self.root:
            print("Binary search tree is empty!")
            return

        status = self._delete_item(item, self.root)
        if status:
            print(f"Successfully deleted {item}!")
        else:
            print(f"{item} not in the tree!")

    def _delete_item(self, item: int, curr: Node | None) -> Node | None:
        if not curr:
            return

        if curr.item == item:
            return self._delete_helper(
                curr,
                curr,
                "root",
            )

        if curr.left_child and curr.left_child.item == item:
            return self._delete_helper(
                curr,
                curr.left_child,
                "left",
            )

        elif curr.right_child and curr.right_child.item == item:
            return self._delete_helper(
                curr,
                curr.right_child,
                "right",
            )

        if curr.item > item:
            return self._delete_item(
                item,
                curr.left_child,
            )
        else:
            return self._delete_item(
                item,
                curr.right_child,
            )

    def _delete_helper(self, parent: Node, to_delete: Node, pos: str) -> Node | None:
        successor = self._find_successor(to_delete.right_child)

        if not successor:
            if to_delete.left_child:
                if pos == "root":
                    self.root = to_delete.left_child
                elif pos == "left":
                    parent.left_child = to_delete.left_child
                else:
                    parent.right_child = to_delete.left_child
            else:
                if pos == "root":
                    self.root = None
                elif pos == "left":
                    parent.left_child = None
                else:
                    parent.right_child = None

            return to_delete

        if (
            to_delete.right_child
            and not to_delete.right_child.left_child
            and to_delete.right_child.item == successor.item
        ):
            to_delete.right_child = to_delete.right_child.right_child

        if pos == "root":
            self.root = successor
        elif pos == "left":
            parent.left_child = successor
        else:
            parent.right_child = successor

        if to_delete.right_child:
            successor.right_child = to_delete.right_child

        if to_delete.left_child:
            successor.left_child = to_delete.left_child

        return to_delete

    def _find_successor(self, curr: Node | None) -> Node | None:
        if not curr:
            return

        current = curr

        while current.left_child:
            if not current.left_child.left_child:
                temp = deepcopy(current)

                if current.left_child.right_child and temp.left_child:
                    current.left_child = temp.left_child.right_child
                else:
                    current.left_child = None

                return temp.left_child

            current = current.left_child

        return current

    def DFS(self) -> None:
        if not self.root:
            print("Binary search tree is empty!\n")

        else:
            path = []
            self._in_order(path, self.root)
            print(path, "\n")

    def _in_order(self, path: list[int], curr: Node | None) -> None:
        if not curr:
            return

        self._in_order(path, curr.left_child)
        path.append(curr.item)
        self._in_order(path, curr.right_child)


def main():
    bst = BST()

    to_insert = [50, 30, 70, 20, 40, 60, 80]
    print(f"Items to insert {to_insert}\n")

    for i in to_insert:
        print(f"Inserting {i}")
        bst.insert(i)
        bst.DFS()

    print(f"Deleting 70")
    bst.delete(70)
    bst.DFS()

    print("Searching 20")
    bst.search(20)
    print()


if __name__ == "__main__":
    main()
