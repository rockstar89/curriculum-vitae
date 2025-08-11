/**
 * Downloads a file from URL
 */
export const downloadFile = (url: string, filename: string): void => {
  const link = document.createElement('a');
  link.href = url;
  link.download = filename;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

/**
 * Opens URL in new tab
 */
export const openInNewTab = (url: string): void => {
  window.open(url, '_blank', 'noopener,noreferrer');
};