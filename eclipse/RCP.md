# How to add new view in the specific folder and make the layout ratio
(Learn from Alex:) Use IPlaceholderFolderLayout:

``` java
public class Perspective implements IPerspectiveFactory {
 
  @Override
  public void createInitialLayout(IPageLayout layout) {
    layout.setEditorAreaVisible(false);
 
    IPlaceholderFolderLayout holderFolder = layout.createPlaceholderFolder(
        "top_view_folder_id", IPageLayout.TOP, 0.7f, IPageLayout.ID_EDITOR_AREA);
    // "top_view" is view primary ID. ":*" is view secondary ID.
    // Which means all new added "top_view" view will be added in this place holder.
    holderFolder.addPlaceholder("top_view" + ":*");  //$NON-NLS-1$
 
    layout.addStandaloneView("bottom_view", true, IPageLayout.BOTTOM, 0.3f, IPageLayout.ID_EDITOR_AREA);
  }
}
```
